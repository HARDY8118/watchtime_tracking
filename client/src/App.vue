<script>

import * as Fingerprint from "@fingerprintjs/fingerprintjs";
import 'vue3-carousel/dist/carousel.css';
import { Carousel, Slide, Pagination, Navigation } from 'vue3-carousel';


export default {
  data() {
    return {
      status: "loading",
      videos: [],
      currentSlide: 0,
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
  },
  async created() {
    this.videos = (await (await fetch("http://localhost:8080/list/videos")).json())
  },
  components: {
    Carousel,
    Slide,
    Pagination,
    Navigation,
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
    <div style="display: flex; justify-content: center;">
      <!-- <b-carousel id="carousel-fade" style="text-shadow: 0px 0px 2px #000; height: 100vh;" fade controls indicators
        img-width="1024" img-height="840">
        <b-carousel-slide :caption="v.title" v-for="(v, i) in videos" :key="i" style="height: 100vh;">
          <video>
            <source :src="'http://127.0.0.1:8080/videos/' + v.file">
          </video>
          <h2>{{ v.title }}</h2>
          <h3>By: {{ v.videographer }}</h3>
        </b-carousel-slide>
      </b-carousel> -->

      <carousel style="width: 90vw; height: 100vh;" :itemsToShow="1" :wrapAround="true">
        <slide v-for="({ videographer, title, file }, i) in videos" :key="i"
          style="display: flex; flex-direction: column;">
          <h2>{{ title }}</h2>
          <br>
          <video autoplay muted style="height: 40vw;">
            <source :src="'http://127.0.0.1:8080/videos/' + file">
          </video>
          <br>
          <h3>By: {{ videographer }}</h3>
        </slide>

        <template #addons="{ slidesCount, currentSlide }">
          <navigation v-if="slidesCount > 1" />
          <!-- <pagination /> -->
          <!-- {{ currentSlide }} -->
          <span ref="current" :current="currentSlide"></span>
        </template>
      </carousel>
    </div>
  </b-overlay>
  <h1>Slide: {{ currentSlide }} </h1>

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
