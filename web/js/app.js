const {createApp, ref} = Vue;

createApp({
  data() {
    return {
      // setup
      tournamentName: '',
      eventName: '',
      eventPhase: '',
      kickertoolLiveURL: '',
      kickertoolLiveURLClass: '',
      discordWebhookURL: '',
      discordWebhookURLClass: '',
      feishuWebhookURL: '',
      feishuWebhookURLClass: '',
      locales: ['en-US', 'zh-CN', 'zh-HK', 'zh-TW', 'ja-JP'],
      selectedLocale: 'en-US',

      // voices
      text: '',
      voices: [],
      selectedVoice: null,

      // matches
      loading: false,
      loadingError: '',
      matches: [],

      // logs
      logs: [],
    };
  },
  mounted() {
    this.loadVoices();
    window.speechSynthesis.onvoiceschanged = this.loadVoices;
  },
  methods: {
    showError(...messages) {
      const formattedMessages = messages.join(' ');
      this.loadingError = formattedMessages;
      this.log('ERROR', formattedMessages);
    },
    showWarn(...messages) {
      const formattedMessages = messages.join(' ');
      this.loadingError = formattedMessages;
      this.log('WARN', formattedMessages);
    },
    log(level = 'INFO', ...messages) {
      const timestamp = new Date().toLocaleTimeString();
      const formattedMessages = messages.join(' ');
      const fullLog = `[${timestamp}] [${level}] ${formattedMessages}`;
      this.logs.unshift(fullLog);
      console.log(fullLog);
    },
    clear() {
      this.logs = [];
    },
    loadVoices() {
      this.voices = window.speechSynthesis.getVoices();
      if (this.voices.length > 0 && !this.selectedVoice) {
        this.selectedVoice = this.voices[0].name; // Select the first available voice
      }
      this.log('INFO', 'Loaded', this.voices.length, 'voice synthesizers');
    },
    updateLocale() {
      if (this.selectedLocale) {
        for (i = 0; i < this.voices.length; i++) {
          if (
            this.voices[i].name.startsWith('Google') &&
            this.selectedLocale == this.voices[i].lang
          ) {
            this.selectedVoice = this.voices[i].name;
            return;
          }
        }
      }
    },
    textToSpeech(text) {
      if (text) {
        const utterance = new SpeechSynthesisUtterance(text);
        const selected = this.voices.find(
          voice => voice.name === this.selectedVoice,
        );
        if (selected) {
          utterance.voice = selected;
        }
        this.log('INFO', 'Spoke [', text, '] with', utterance.voice.name);
        window.speechSynthesis.speak(utterance);
      }
    },
    async buildSpeakText(match, provider, template) {
      try {
        this.loadingError = 'Sending...';
        const url = '/notify';
        const params = {
          tournamentName: this.tournamentName,
          eventName: this.eventName,
          eventPhase: this.eventPhase,
          team1: match.team1,
          team2: match.team2,
          tableNo: match.tableNo,
          locale: this.selectedLocale,
          msgType: provider,
          template: template,
        };
        const response = await fetch(url, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(params),
        });
        if (!response.ok) {
          this.showError('Network response was not ok:', response.statusText);
        }
        this.loadingError = '';
        const data = await response.json();
        const text = data.data.text;
        return text;
      } catch (err) {
        this.showError(err.message);
      } finally {
        this.loading = false;
      }
      return '';
    },
    async notify(match, provider, template) {
      try {
        this.loadingError = 'Sending...';
        const url = '/notify';
        const params = {
          tournamentName: this.tournamentName,
          eventName: this.eventName,
          eventPhase: this.eventPhase,
          team1: match.team1,
          team2: match.team2,
          tableNo: match.tableNo,
          locale: this.selectedLocale,
          msgType: provider,
          template: template,
        };
        if (provider == 'discord') {
          params.discordWebhookURL = this.discordWebhookURL;
        } else if (provider == 'feishu') {
          params.feishuWebhookURL = this.feishuWebhookURL;
        } else {
          this.showError('Illegal provider:', provider);
          return;
        }
        const response = await fetch(url, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(params),
        });
        if (!response.ok) {
          this.showError('Network response was not ok:', response.statusText);
        }
        this.loadingError = '';
        const data = await response.json();
        this.log(
          'INFO',
          'Notified match provider:',
          provider,
          'template:',
          template,
          'text:',
          data.data.text,
        );
      } catch (err) {
        this.showError(err.message);
      } finally {
        this.loading = false;
      }
    },
    currentMatch(index) {
      const match = JSON.parse(JSON.stringify(this.matches[index]));
      match.team1 = match.team1.map(player => player.id);
      match.team2 = match.team2.map(player => player.id);
      return match;
    },
    async handleCall(index) {
      this.log('INFO', 'Called match index:', index);
      const match = this.currentMatch(index);
      // speak
      this.textToSpeech(
        await this.buildSpeakText(match, 'speak', 'call_match'),
      );
      // web hooks
      const template = 'call_match';
      if (this.discordWebhookURL) {
        await this.notify(match, 'discord', template);
      }
      if (this.feishuWebhookURL) {
        await this.notify(match, 'feishu', template);
      }
    },
    async handleEdit(index) {
      this.log('INFO', 'Edited match index:', index);
      const match = this.currentMatch(index);
      this.text = await this.buildSpeakText(match, 'speak', 'call_match');
    },
    async handleRecall(index, id) {
      this.log('INFO', 'Recalled player ID:', id);
      const match = JSON.parse(JSON.stringify(this.matches[index]));
      match.team1 = [id];
      match.team2 = [];
      // speak
      this.textToSpeech(
        await this.buildSpeakText(match, 'speak', 'recall_player'),
      );
      // web hooks
      const template = 'recall_player';
      if (this.discordWebhookURL) {
        await this.notify(match, 'discord', template);
      }
      if (this.feishuWebhookURL) {
        await this.notify(match, 'feishu', template);
      }
    },
    async handleAnnounce() {
      if (this.text) {
        this.log('INFO', 'Announced text:', this.text);
        this.textToSpeech(this.text);
        if (this.discordWebhookURL) {
          this.log('INFO', 'Sent text to Discord:', this.text);
          await this.notifyManually('discord', this.text);
        }
        if (this.feishuWebhookURL) {
          this.log('INFO', 'Sent text to Feishu:', this.text);
          await this.notifyManually('feishu', this.text);
        }
      } else {
        this.showWarn('Please input texts manually.');
      }
    },
    async notifyManually(provider, text) {
      try {
        this.loadingError = 'Sending...';
        const url = '/notify';
        const params = {
          msgType: provider,
          text: text,
        };
        if (provider == 'discord') {
          params.discordWebhookURL = this.discordWebhookURL;
        } else if (provider == 'feishu') {
          params.feishuWebhookURL = this.feishuWebhookURL;
        } else {
          this.showError('Illegal provider:', provider);
          return;
        }
        const response = await fetch(url, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(params),
        });
        if (!response.ok) {
          this.showError('Network response was not ok:', response.statusText);
        }
        this.loadingError = '';
      } catch (err) {
        this.showError(err.message);
      } finally {
        this.loading = false;
      }
    },
    validateFeishuWebhookURL() {
      if (
        !(
          this.feishuWebhookURL.startsWith(
            'https://open.feishu.cn/open-apis/bot/v2/hook/',
          ) ||
          this.feishuWebhookURL.startsWith(
            'https://open.larkoffice.com/open-apis/bot/v2/hook/',
          ) ||
          this.feishuWebhookURL.startsWith(
            'https://open.larksuite.com/open-apis/bot/v2/hook/',
          )
        )
      ) {
        this.feishuWebhookURLClass = 'panel-item-error';
        return false;
      }

      this.feishuWebhookURLClass = '';
      return true;
    },
    validateDiscordWebhookURL() {
      if (
        !this.discordWebhookURL.startsWith('https://discord.com/api/webhooks/')
      ) {
        this.discordWebhookURLClass = 'panel-item-error';
        return false;
      }

      this.discordWebhookURLClass = '';
      return true;
    },
    validateKickertoolLiveURL() {
      if (!this.kickertoolLiveURL) {
        this.kickertoolLiveURLClass = 'panel-item-error';
        return false;
      }
      if (
        !(
          this.kickertoolLiveURL.startsWith('https://live.kickertool.de') &&
          this.kickertoolLiveURL.endsWith('/live')
        )
      ) {
        this.kickertoolLiveURLClass = 'panel-item-error';
        return false;
      }

      this.kickertoolLiveURLClass = '';
      return true;
    },
    async handleCrawl() {
      if (!this.validateKickertoolLiveURL()) {
        this.showWarn('Kickertool Live URL is not valid');
        return;
      }
      this.log('INFO', 'Crawled', this.kickertoolLiveURL);
      try {
        this.loadingError = 'Loading...';
        const url = '/crawl?url=' + this.kickertoolLiveURL;
        const response = await fetch(url);
        if (!response.ok) {
          this.showError('Network response was not ok:', response.statusText);
        }
        const data = await response.json();
        this.matches = data.data.matches;
        this.loadingError = '';
      } catch (err) {
        this.showError(err.message);
      } finally {
        this.loading = false;
      }
    },
  },
}).mount('#app');
