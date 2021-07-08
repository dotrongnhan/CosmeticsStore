<template>
  <div>
    <!-- Slide1 -->
    <section class="slide1">
      <div class="wrap-slick1">
        <Carousel :wrapAround="true">
          <Slide v-for="slide in slides" :key="slide.id">
            <div
              class="item-slick1"
              :style="{
                backgroundImage: 'url(' + slide.image + ')',
                width: '100%',
              }"
            >
              <div
                class="
                  wrap-content-slide1
                  sizefull
                  flex-col-c-m
                  p-l-15 p-r-15 p-t-200 p-b-100
                "
              >
                <h5 style="font-family: emoji; color: white" class="caption2-slide1 l-text1 t-center animated m-b-37">
                  For women, never miss an opportunity to tell them they’re beautiful
                </h5>

                <div class="wrap-btn-slide1 w-size1 animated">
                  <router-link
                    to="/products"
                    class="
                      flex-c-m
                      size2
                      bo-rad-23
                      s-text2
                      bgwhite
                      hov1
                      trans-0-4
                    "
                  >
                    Shop Now
                  </router-link>
                </div>
              </div>
            </div>
          </Slide>
        </Carousel>
      </div>
    </section>

    <SectionCategory />
    <SectionBestSale />

    <!-- Shipping -->
    <SectionShipping />
  </div>
</template>

<script>
import "vue3-carousel/dist/carousel.css";
import { Carousel, Slide } from "vue3-carousel";
import { mapActions } from "vuex";
import SectionCategory from "./SectionCategory.vue";
import SectionBestSale from "./SectionBestSale.vue";
import SectionShipping from "./SectionShipping.vue";

export default {
  name: "HomePage",

  components: {
    Carousel,
    Slide,
    SectionCategory,
    SectionBestSale,
    SectionShipping,
  },

  data() {
    return {
      slides: [
        { id: 1, image: require("@/assets/silde1.jpeg") },
        { id: 2, image: require("@/assets/slide2.jpg") },
        { id: 3, image: require("@/assets/slide3.jpg") },
      ],
    };
  },
  created() {
    this.$store.dispatch("products/getProducts");
  },
  methods: {
    ...mapActions(["products/getProducts"]),
  },
};
</script>