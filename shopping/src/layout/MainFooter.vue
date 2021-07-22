<template>
  <footer class="bg6 p-t-45 p-b-45 p-l-45 p-r-45">
    <div class="flex-w">
      <div class="w-size6 p-t-30 p-l-15 p-r-15 respon3">
        <h4 class="s-text12 p-b-30">ADDRESS</h4>

        <div>
          <p class="s-text7 w-size27">
          Onoxa 2260 118th Ave N St. Petersburg FL 33716
          </p>

          <h4 class="s-text12 p-b-30">SOCIAL HUB</h4>
          <div class="flex-m p-t-30" style="margin-top: -30px;">
            <a href="https://www.facebook.com/onoxa/" class="fs-18 color1 p-r-20 fa fa-facebook"></a>
            <a href="https://www.instagram.com/onoxa_com/" class="fs-18 color1 p-r-20 fa fa-instagram"></a>
            <a href="https://www.linkedin.com/company/onoxa" class="fs-18 color1 p-r-20 fa fa-linkedin"></a>
          </div>
        </div>
      </div>

      <div class="w-size7 p-t-30 p-l-15 p-r-15 respon4">
        <h4 class="s-text12 p-b-30  text-center">Categories</h4>

        <ul>
          <li class="p-b-9" @click="summitCategory({ value: 'Toner', label: 'Toner' })">
            <router-link to="/products" class="s-text7" > Toner </router-link>
          </li>

          <li class="p-b-9" @click="summitCategory({ value: 'Lipstick', label: 'Lipstick' })">
            <router-link to="/products" class="s-text7" > Lipstick </router-link>
          </li>

          <li class="p-b-9" @click="summitCategory({ value: 'Face Scream', label: 'Face Scream' })">
            <router-link to="/products" class="s-text7" > Face Scream </router-link>
          </li>

          <li class="p-b-9" @click="summitCategory({ value: 'Mask', label: 'Mask' })">
            <router-link to="/products" class="s-text7" > Mask </router-link>
          </li>

          <li class="p-b-9" @click="summitCategory({ value: 'Serum', label: 'Serum' })">
            <router-link to="/products" class="s-text7" > Serum </router-link>
          </li>
        </ul>
      </div>

      <div class="w-size7 p-t-30 p-l-15 p-r-15 respon4">
        <h4 class="s-text12 p-b-30 text-center">Links</h4>

        <ul>
          <li class="p-b-9">
            <a href="#" class="s-text7"> Home </a>
          </li>

          <li class="p-b-9">
            <a href="/about" class="s-text7"> Products </a>
          </li>

          <li class="p-b-9">
            <a href="/contact" class="s-text7"> About </a>
          </li>

          <li class="p-b-9">
            <a href="/" class="s-text7"> Contact </a>
          </li>
        </ul>
      </div>

      <div class="w-size7 p-t-30 p-l-15 p-r-15 respon4">
        <h4 class="s-text12 p-b-30  text-center">Help</h4>

        <ul>
          <li class="p-b-9">
            <a href="#" class="s-text7"> Facebook </a>
          </li>

          <li class="p-b-9">
            <a href="#" class="s-text7"> instagram </a>
          </li>

          <li class="p-b-9">
            <a href="#" class="s-text7"> About us </a>
          </li>
        </ul>
      </div>

      <div class="w-size8 p-t-30 p-l-15 p-r-15 respon3">
        <h4 class="s-text12 p-b-30">Newsletter</h4>

        <form>
          <div class="effect1 w-size9">
            <input
              class="s-text7 bg6 w-full p-b-5"
              type="text"
              name="email"
              placeholder="email@example.com"
            />
            <span class="effect1-line"></span>
          </div>

          <div class="w-size2 p-t-20">
            <!-- Button -->
            <button class="flex-c-m size2 bg4 bo-rad-23 hov1 m-text3 trans-0-4">
              Subscribe
            </button>
          </div>
        </form>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  name: "MainFooter",
  methods: {
    summitCategory (category) {
          this.$store.commit("products/GET_PRODUCT_BY_CATEGORY", category)
          this.$store.dispatch("products/getProducts", {offset: 1, sortType: 'Default', category: category.value})
    }
  },
};
</script>
