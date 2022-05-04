import { createApp } from "vue";
import App from "./App.vue";
import BootstrapVue3 from "bootstrap-vue-3";
import "fullpage.js/vendors/scrolloverflow";
import VueFullPage from "vue-fullpage.js";

import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue-3/dist/bootstrap-vue-3.css";

createApp(App).use(BootstrapVue3).use(VueFullPage).mount("#app");
