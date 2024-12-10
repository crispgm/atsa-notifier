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
      locales: ['enUS', 'zhCN'],
      selectedLocale: 'enUS',

      // voices
      text: '',
      voices: [],
      selectedVoice: null,

      // matches
      loading: false,
      loadingError: '',
      matches: [],

      // logs
    };
  },
  mounted() {
    this.loadVoices();
    window.speechSynthesis.onvoiceschanged = this.loadVoices;
  },
  methods: {
    loadVoices() {
      this.voices = window.speechSynthesis.getVoices();
      if (this.voices.length > 0 && !this.selectedVoice) {
        this.selectedVoice = this.voices[0].name; // Select the first available voice
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
        window.speechSynthesis.speak(utterance);
      }
    },
    speakText() {
      if (this.text) {
        this.textToSpeech(this.text)
      }
    },
    async buildMatchText(index) {
      try {
        this.loadingError = 'sending';
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
        };
        const response = await fetch(url, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(params),
        });
        if (!response.ok) {
          throw new Error(
            'Network response was not ok: ' + response.statusText,
          );
        }
        this.loadingError = '';
        const data = await response.json();
        const text = data.data.text;
        return text
      } catch (err) {
        this.loadingError = err.message;
      } finally {
        this.loading = false;
      }
      return '';
    },
    async call(index) {
      this.textToSpeech(await this.buildMatchText(index));
    },
    async edit(index) {
      this.text = await this.buildMatchText(index);
    },
    async notifyText() {
      if (!this.discordWebhookURL) {
        console.log('discordWebhookURL is not set');
        return
      }
      try {
        this.loadingError = 'sending';
        const url = '/notify';
        const params = {
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
          throw new Error(
            'Network response was not ok: ' + response.statusText,
          );
        }
        this.loadingError = '';
      } catch (err) {
        this.loadingError = err.message;
      } finally {
        this.loading = false;
      }
    },
    async notify(index) {
      if (!this.discordWebhookURL) {
        console.log('discordWebhookURL is not set');
        return
      }
      try {
        this.loadingError = 'sending';
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
        };
        const response = await fetch(url, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(params),
        });
        if (!response.ok) {
          throw new Error(
            'Network response was not ok: ' + response.statusText,
          );
        }
        this.loadingError = '';
      } catch (err) {
        this.loadingError = err.message;
      } finally {
        this.loading = false;
      }
    },
    async crawl() {
      if (!this.kickertoolLiveURL) {
        console.log('kickertoolLiveURL is not set');
        return
      }
      try {
        this.loadingError = 'loading';
        const url = "/crawl?url=" + this.kickertoolLiveURL;
        const response = await fetch(url);
        if (!response.ok) {
          throw new Error(
            'Network response was not ok: ' + response.statusText,
          );
        }
        const data = await response.json();
        this.matches = data.data.matches;
        this.loadingError = '';
      } catch (err) {
        this.loadingError = err.message;
      } finally {
        this.loading = false;
      }
    },
  },
}).mount('#app');
