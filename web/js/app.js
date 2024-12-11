const {createApp, ref} = Vue;

createApp({
  data() {
    return {
      // setup
      tournamentName: '',
      eventName: '',
      eventPhase: '',
      kickertoolLiveURL: '',
      discordWebhookURL: '',
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
    showError(msg) {
      this.loadingError = msg;
      this.log('ERROR', msg)
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
          if (this.voices[i].name.startsWith('Google') && this.selectedLocale == this.voices[i].lang) {
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
        this.log('INFO', 'Spoke [' + text + '] with', utterance.voice.name);
        window.speechSynthesis.speak(utterance);
      }
    },
    async buildMatchText(type, template, mIndex, tIndex = -1, pIndex = -1) {
      try {
        this.loadingError = 'Sending...';
        const url = '/notify';
        const match = this.matches[mIndex];
        const params = {
          tournamentName: this.tournamentName,
          eventName: this.eventName,
          eventPhase: this.eventPhase,
          team1: match.team1,
          team2: match.team2,
          tableNo: match.tableNo,
          locale: this.selectedLocale,
          msgType: type,
          template: template,
        };
        if (template == 'recall_player') {
          let team = match.team1
          if (tIndex == 1) {
            team = match.team2
          }
          console.log(team[pIndex])
          params.team1 = [team[pIndex]];
          params.team2 = [];
        }
        const response = await fetch(url, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(params),
        });
        if (!response.ok) {
          this.showError('Network response was not ok: ' + response.statusText);
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
    async call(index) {
      this.textToSpeech(await this.buildMatchText('speak', 'call_match', index));
      this.log('INFO', 'Called match index:', index);
    },
    async edit(index) {
      this.text = await this.buildMatchText('speak', 'call_match', index);
      this.log('INFO', 'Edited match index:', index);
    },
    async recall(mIndex, tIndex, pIndex) {
      this.textToSpeech(await this.buildMatchText('speak', 'recall_player', mIndex, tIndex, pIndex));
      this.log('INFO', 'Called match index:', mIndex, tIndex, pIndex);
    },
    speakText() {
      if (this.text) {
        this.textToSpeech(this.text);
      } else {
        this.showError('Please input texts manually.')
      }
    },
    async notifyText() {
      if (!this.discordWebhookURL) {
        this.showError('discordWebhookURL is not set');
        return;
      }
      try {
        this.loadingError = 'Sending...';
        const url = '/notify';
        const params = {
          msgType: 'discord',
          text: this.text,
          discordWebhookURL: this.discordWebhookURL,
        };
        const response = await fetch(url, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(params),
        });
        if (!response.ok) {
          this.showError('Network response was not ok: ' + response.statusText);
        }
        this.loadingError = '';
      } catch (err) {
        this.showError(err.message);
      } finally {
        this.loading = false;
      }
    },
    async notify(index) {
      if (!this.discordWebhookURL) {
        this.showError('discordWebhookURL is not set');
        return;
      }
      try {
        this.loadingError = 'Sending...';
        const url = '/notify';
        const match = this.matches[index];
        const params = {
          tournamentName: this.tournamentName,
          eventName: this.eventName,
          eventPhase: this.eventPhase,
          team1: match.team1,
          team2: match.team2,
          tableNo: match.tableNo,
          locale: this.selectedLocale,
          discordWebhookURL: this.discordWebhookURL,
          msgType: 'discord',
          template: 'call_match',
        };
        const response = await fetch(url, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(params),
        });
        if (!response.ok) {
          this.showError('Network response was not ok: ' + response.statusText);
        }
        this.loadingError = '';
        const data = await response.json();
        this.log('INFO', 'Notified match index [', index, '] text=', data.data.text);
      } catch (err) {
        this.showError(err.message);
      } finally {
        this.loading = false;
      }
    },
    async crawl() {
      if (!this.kickertoolLiveURL) {
        this.showError('kickertoolLiveURL is not set');
        return;
      }
      this.log('INFO', 'Crawled', this.kickertoolLiveURL);
      try {
        this.loadingError = 'Loading...';
        const url = '/crawl?url=' + this.kickertoolLiveURL;
        const response = await fetch(url);
        if (!response.ok) {
          this.showError('Network response was not ok: ' + response.statusText);
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
