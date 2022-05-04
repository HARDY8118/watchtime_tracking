<script>

import * as Fingerprint from "@fingerprintjs/fingerprintjs";

export default {
  data() {
    return {
      status: "loading",
      videos: []
    }
  },
  methods: {
    onSlideStart(slide) {
      this.sliding = true
    },
    onSlideEnd(slide) {
      this.sliding = false
    }
  },
  async mounted() {
    this.status = "IDing"
    // await new Promise(resolve => setTimeout(resolve, 3000))
    try {
      const result = await (await Fingerprint.load()).get();
      console.log(result.visitorId)
      // await new Promise(resolve => setTimeout(resolve, 5000))
      this.status = "Done"
    } catch (e) {
      console.error(e)
    }

    // document.querySelector('.carousel-item').classList.add('active')
  },
  async created() {
    this.videos = (await (await fetch("http://localhost:8080/list/videos")).json())
  }
}
</script>

<template>
  <b-overlay :show="status !== 'Done'" rounded="sm">
    <template #overlay>
      <b-alert show variant="info" style="width: 90vw; text-align: center;">
        <span class="overlay-status">{{ status }}</span>
        <br>
      </b-alert>
    </template>
    <div>
      <b-carousel id="carousel-fade" style="text-shadow: 0px 0px 2px #000; height: 100vh;" fade controls indicators
        img-width="1024" img-height="840">
        <b-carousel-slide :caption="v.title" v-for="(v, i) in videos" :key="i" style="height: 100vh;">
          <video>
            <source :src="'http://127.0.0.1:8080/videos/' + v.file">
          </video>
          <h2>{{ v.title }}</h2>
          <h3>By: {{ v.videographer }}</h3>
        </b-carousel-slide>
      </b-carousel>
    </div>
  </b-overlay>

  <!-- <video controls>
    <source src="http://localhost:8080/videos/pexels-вальдемар-10026607.mp4">
  </video> -->

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
