<template>
    <vueper-slides @ready="slideStart($event)" @next="slideEnd($event)" @previous="slideEnd($event)"
        @slide="slideStart($event)" ref="slider" fade :touchable="false" :bullets="false" :arrows="true">
        <vueper-slide v-for="slide in images" :key="slide.id"
            :image="config.serverAddress + '/images/' + slide.filename" :title="slide.title"
            :content="'By ' + slide.photographer" />

        <template #arrow-left>
            <svg viewBox="0 0 9 18">
                <path stroke-linecap="round" d="m8 1 l-7 8 7 8"></path>
            </svg>
        </template>

        <template #arrow-right>
            <svg viewBox="0 0 9 18">
                <path stroke-linecap="round" d="m1 1 l7 8 -7 8"></path>
            </svg>
        </template>

    </vueper-slides>

</template>
<script>

import 'vueperslides/dist/vueperslides.css'
import { VueperSlides, VueperSlide } from 'vueperslides'
import config from "../assets/config.json"

export default {
    name: 'Slider',
    components: { VueperSlides, VueperSlide },
    data() {
        return {
            slides: [],
            config,
            currentImageId: null
        }
    },
    props: {
        images: Array,
        fingerprintid: String
    },
    methods: {
        slideEnd({ currentSlide }) {
            if (currentSlide) {
                console.log("ended", currentSlide.image.match(/(\d+)\.jpg/)[1], Date.now())
                navigator.sendBeacon(config.serverAddress + "/switch", `${this.$props.fingerprintid}:ended:${currentSlide.image.match(/(\d+)\.jpg/)[1]}:${Date.now()}`)
            }
        },
        slideStart({ currentSlide }) {
            if (currentSlide) {
                console.log("started", currentSlide.image.match(/(\d+)\.jpg/)[1], Date.now())
                navigator.sendBeacon(config.serverAddress + "/switch", `${this.$props.fingerprintid}:started:${currentSlide.image.match(/(\d+)\.jpg/)[1]}:${Date.now()}`)
            }
        },
    },
    watch: {
        images(newVal, oldVal) {
            if (!!newVal) {
                // console.log("Started ", this.currentImageId, Date.now())
            }
        }
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