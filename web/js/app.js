const {createApp, ref} = Vue;

createApp({
  data() {
    return {
      // setup
      setup: {
        tournamentName: '',
        eventName: '',
        eventPhase: '',
        text: '',
        selectedLocale: 'en-US',
        selectedVoice: null,
      },

      kickertoolLiveURL: '',
      kickertoolLiveURLClass: '',
      discordWebhookURL: '',
      discordWebhookURLClass: '',
      feishuWebhookURL: '',
      feishuWebhookURLClass: '',
      locales: ['en-US', 'zh-CN', 'zh-HK', 'zh-TW', 'ja-JP'],
      voices: [],

      loading: false,
      loadingError: '',
      // matches
      matches: [],
      // logs
      logs: [],
    };
  },
  created() {
    const setup = localStorage.getItem('setup');
    if (setup) {
      this.setup = JSON.parse(setup);
    }
    const kickertoolLiveURL = localStorage.getItem('kickertoolLiveURL');
    if (kickertoolLiveURL) {
      this.kickertoolLiveURL = kickertoolLiveURL;
    }
    const discordWebhookURL = localStorage.getItem('discordWebhookURL');
    if (discordWebhookURL) {
      this.discordWebhookURL = discordWebhookURL;
    }
    const feishuWebhookURL = localStorage.getItem('feishuWebhookURL');
    if (feishuWebhookURL) {
      this.feishuWebhookURL = feishuWebhookURL;
    }
  },
  mounted() {
    this.loadVoices();
    window.speechSynthesis.onvoiceschanged = this.loadVoices;
  },
  watch: {
    setup: {
      handler(newData) {
        localStorage.setItem('setup', JSON.stringify(newData));
      },
      deep: true,
    },
    kickertoolLiveURL: {
      handler(newData) {
        if (!this.kickertoolLiveURL || this.validateKickertoolLiveURL()) {
          localStorage.setItem('kickertoolLiveURL', newData);
        }
      },
      deep: true,
    },
    discordWebhookURL: {
      handler(newData) {
        if (!this.discordWebhookURL || this.validateDiscordWebhookURL()) {
          localStorage.setItem('discordWebhookURL', newData);
        }
      },
      deep: true,
    },
    feishuWebhookURL: {
      handler(newData) {
        if (!this.feishuWebhookURL || this.validateFeishuWebhookURL()) {
          localStorage.setItem('feishuWebhookURL', newData);
        }
      },
      deep: true,
    },
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
    handleClear() {
      this.logs = [];
    },
    handleReset() {
      this.handleClear();
      this.setup = {
        tournamentName: '',
        eventName: '',
        eventPhase: '',
        text: '',
        selectedLocale: 'en-US',
        selectedVoice: null,
      };
      this.kickertoolLiveURL = '';
      this.discordWebhookURL = '';
      this.feishuWebhookURL = '';
    },
    loadVoices() {
      this.voices = window.speechSynthesis.getVoices();
      if (this.voices.length > 0 && !this.setup.selectedVoice) {
        this.setup.selectedVoice = this.voices[0].name; // Select the first available voice
      }
    },
    updateLocale() {
      if (this.setup.selectedLocale) {
        for (i = 0; i < this.voices.length; i++) {
          if (
            this.voices[i].name.startsWith('Google') &&
            this.setup.selectedLocale == this.voices[i].lang
          ) {
            this.setup.selectedVoice = this.voices[i].name;
            return;
          }
        }
      }
    },
    textToSpeech(text) {
      if (text) {
        const utterance = new SpeechSynthesisUtterance(text);
        const selected = this.voices.find(
          voice => voice.name === this.setup.selectedVoice,
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
          tournamentName: this.setup.tournamentName,
          eventName: this.setup.eventName,
          eventPhase: this.setup.eventPhase,
          locale: this.setup.selectedLocale,
          team1: match.team1,
          team2: match.team2,
          tableNo: match.tableNo,
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
          tournamentName: this.setup.tournamentName,
          eventName: this.setup.eventName,
          eventPhase: this.setup.eventPhase,
          team1: match.team1,
          team2: match.team2,
          tableNo: match.tableNo,
          locale: this.setup.selectedLocale,
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
    async handleCall(match) {
      this.log('INFO', 'Called match at table:', match.tableNo);
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
    async handleEdit(match) {
      this.log('INFO', 'Edited match at table:', match.tableNo);
      this.setup.text = await this.buildSpeakText(match, 'speak', 'call_match');
    },
    async handleRecall(match, player) {
      this.log('INFO', 'Recalled player:', player.name);
      const tempMatch = {
        tableNo: match.tableNo,
        team1: [player],
        team2: [],
      };
      // speak
      this.textToSpeech(
        await this.buildSpeakText(tempMatch, 'speak', 'recall_player'),
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
      if (this.setup.text) {
        this.log('INFO', 'Announced text:', this.setup.text);
        this.textToSpeech(this.setup.text);
        if (this.discordWebhookURL) {
          this.log('INFO', 'Sent text to Discord:', this.setup.text);
          await this.notifyManually('discord', this.setup.text);
        }
        if (this.feishuWebhookURL) {
          this.log('INFO', 'Sent text to Feishu:', this.setup.text);
          await this.notifyManually('feishu', this.setup.text);
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
      if (this.feishuWebhookURL) {
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
      }

      this.feishuWebhookURLClass = '';
      return true;
    },
    validateDiscordWebhookURL() {
      if (this.discordWebhookURL) {
        if (
          !this.discordWebhookURL.startsWith(
            'https://discord.com/api/webhooks/',
          )
        ) {
          this.discordWebhookURLClass = 'panel-item-error';
          return false;
        }
      }

      this.discordWebhookURLClass = '';
      return true;
    },
    validateKickertoolLiveURL() {
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
    async handleSync() {
      if (!this.kickertoolLiveURL) {
        this.kickertoolLiveURLClass = 'panel-item-error';
        this.showWarn('Kickertool Live URL is not set');
        return;
      }
      if (!this.validateKickertoolLiveURL()) {
        this.showWarn('Kickertool Live URL is not valid');
        return;
      }
      try {
        this.loadingError = 'Loading...';
        const url = '/sync?url=' + this.kickertoolLiveURL;
        const response = await fetch(url);
        if (!response.ok) {
          this.showError('Network response was not ok:', response.statusText);
        }
        const data = await response.json();
        this.matches = data.data.matches;
        if (!this.matches) {
          this.matches = [];
        }
        this.log('INFO', 'Synced', this.matches.length, 'matches from', this.kickertoolLiveURL);
        this.loadingError = '';
      } catch (err) {
        this.showError(err.message);
      } finally {
        this.loading = false;
      }
    },
  },
}).mount('#app');
