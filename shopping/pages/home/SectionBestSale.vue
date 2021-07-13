<template>
  <section class="instagram p-t-20">
      <div class="sec-title p-b-52 p-l-15 p-r-15">
        <h3 class="m-text5 t-center">Our best sale</h3>
      </div>
      <div class="flex-w">
        <div v-for="(product, index) in products.products.slice(0, 5)" :key="index" class="block4 wrap-pic-w">
          <img :src="product.image" alt="IMG-INSTAGRAM" />

          <router-link :to="'/products/' + product.id" class="block4-overlay sizefull ab-t-l trans-0-4">
            <span
              class="
                block4-overlay-heart
                s-text9
                flex-m
                trans-0-4
                p-l-40 p-t-25
              "
            >
              <i class="icon_heart_alt fs-20 p-r-12" aria-hidden="true"></i>
              <span class="p-t-2">{{product.product_name}}</span>
            </span>

            <div class="block4-overlay-txt trans-0-4 p-l-40 p-r-25 p-b-30">
              <p class="s-text10 m-b-15 h-size1 of-hidden">
                {{product.description}}
              </p>

              <span class="s-text9"> $ {{product.price}}</span>
            </div>
          </router-link>
        </div>
      </div>
    </section>
</template>

<script>
import { mapState} from "vuex";

export default {
  name: "SectionBestSale",
  props: ['product'],

  computed: mapState(["products"])
}
</script>