<script>

import * as Fingerprint from "@fingerprintjs/fingerprintjs";
import Slider from "./components/Slider.vue";
import config from "./assets/config.json"

export default {
  data() {
    return {
      status: "loading",
      images: [],
      fingerprintid: null,
      config
    }
  },
  methods: {
  },
  async mounted() {
    this.status = "IDing"
    // await new Promise(resolve => setTimeout(resolve, 3000))
    try {
      const result = await (await Fingerprint.load()).get();
      // console.log(result.visitorId)
      this.fingerprintid = result.visitorId
      // await new Promise(resolve => setTimeout(resolve, 5000))
      // this.status = "Done"
    } catch (e) {
      console.error(e)
    }
  },
  async created() {
    this.images = (await (await fetch(config.serverAddress + "/list/images")).json())
  },
  components: { Slider }
}
</script>

<template>
  <b-overlay :show="!fingerprintid" rounded="sm">
    <template #overlay>
      <b-alert show variant="info" style="width: 90vw; text-align: center;">
        <span class="overlay-status">{{ status }}</span>
        <br>
      </b-alert>
    </template>
    <Slider :images="images" :fingerprintid="fingerprintid" />
  </b-overlay>
</template>

<style>
#app {
  height: 100vh;
}

@keyframes loadingdots {
  0% {
    content: "";
  }

  25% {
    content: ".";
  }

  50% {
    content: "..";
  }

  75% {
    content: "...";
  }

  90% {
    content: "....";
  }

  100% {
    content: "";
  }
}

.overlay-status::after {
  animation: loadingdots 1s ease infinite;
  content: "";
}
</style>
