<template>
    <vueper-slides fade :touchable="false" :bullets="false">
        <vueper-slide v-for="(slide, i) in images" :key="i" :image="'http://127.0.0.1:8080/images/' + slide.filename"
            :title="slide.title" :content="'By ' + slide.photographer" />
    </vueper-slides>
</template>

    <!-- <vueper-slides ref="videoSlides">
      <vueper-slide v-for="(v, i) in videos" :key="i" :title="v.title" :content="v.videographer" :video="{
        url: 'http://127.0.0.1:8080/videos/' + v.file, props: { pointerEvents: false, allow: 'autoplay', autoplay: true, loop: true, controls: false, muted: true }
      }"></vueper-slide>
      <template #arrow-left>&lt;PREV</template>
      <template #arrow-right>NEXT&gt;</template>
    </vueper-slides> -->
<script>

import 'vueperslides/dist/vueperslides.css'
import { VueperSlides, VueperSlide } from 'vueperslides'

export default {
    name: 'Slider',
    components: { VueperSlides, VueperSlide },
    data() {
        return {
            slides: []
        }
    },
    props: {
        images: Array
    }
}
</script>

<style>
.vueperslides__parallax-wrapper {
    min-height: 640px;
}

.vueperslides:hover .vueperslides__arrow {
    opacity: 1;
    transform-origin: center;
}

.vueperslide__title {
    font-size: 2em;
    transition: 0.5s all ease-in-out;
    opacity: 0.6;
}

.vueperslides:hover .vueperslide__title {
    font-weight: bold;
    font-size: 2.4em;
    opacity: 0.8;
    background-color: white;
}

.vueperslide__content {
    font-size: 1.8em;
    transition: 0.5s all ease-in-out;
    opacity: 0.6;
}

.vueperslides:hover .vueperslide__content {
    font-weight: bold;
    font-size: 2em;
    opacity: 0.8;
    background-color: gray;
}
</style>