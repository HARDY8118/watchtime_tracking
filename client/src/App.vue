<script>

import * as Fingerprint from "@fingerprintjs/fingerprintjs";

export default {
  data() {
    return {
      status: "loading",
      videos: [],
      currentSlide: 0,
    }
  },
  methods: {
    show() {
      console.log(document.querySelector('.active').getAttribute('data-key'))
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
      <div id="carouselExampleControls" class="carousel slide" data-bs-ride="carousel" data-bs-interval="0">
        <div class="carousel-inner">
          <div class="carousel-item" v-for="(v, i) in videos" :key="i" :class="{ active: i == 0 }" :data-key="i">
            <video autoplay muted loop>
              <source :src="'http://127.0.0.1:8080/videos/' + v.file">
            </video>
            <h2>{{ v.title }}</h2>
            <h3>By: {{ v.videographer }}</h3>
          </div>
        </div>
        <button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleControls"
          data-bs-slide="prev">
          <span class="carousel-control-prev-icon" aria-hidden="true"></span>
          <span class="visually-hidden">Previous</span>
        </button>
        <button class="carousel-control-next" type="button" data-bs-target="#carouselExampleControls"
          data-bs-slide="next">
          <span class="carousel-control-next-icon" aria-hidden="true"></span>
          <span class="visually-hidden">Next</span>
        </button>
      </div>
    </div>
  </b-overlay>
  <h1 @click="show">Slide: {{ currentSlide }} </h1>

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
