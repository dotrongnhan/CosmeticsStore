<template>
  <section class="banner bgwhite p-t-40 p-b-40">
    <div class="container">
      <div class="row">
        <div class="col-sm-10 col-md-8 col-lg-4 m-l-r-auto">
          <!-- block1 -->
          <div class="block1 hov-img-zoom pos-relative m-b-30">
            <img src="https://m.media-amazon.com/images/S/aplus-media/vc/9367a902-f4c8-455e-9e01-23a4051f5d11.__CR0,0,300,300_PT0_SX300_V1___.jpg" alt="IMG-BENNER" />

            <div @click="summitCategory(5)" class="block1-wrapbtn w-size2">
              <router-link to="/products" class="flex-c-m size2 m-text2 bg3 hov1 trans-0-4">
                LIPSTICK
              </router-link>
            </div>
          </div>

          <!-- block1 -->
          <div class="block1 hov-img-zoom pos-relative m-b-30">
            <img src="https://cdn.chanhtuoi.com/viectainha/2021/02/w800/w400/serum-duong-da-cang-bong-silk.jpg.webp" alt="IMG-BENNER" />

            <div @click="summitCategory(6)" class="block1-wrapbtn w-size2">
              <!-- Button -->
              <router-link to="/products"  class="flex-c-m size2 m-text2 bg3 hov1 trans-0-4">
                SERUM
              </router-link>
            </div>
          </div>
        </div>

        <div class="col-sm-10 col-md-8 col-lg-4 m-l-r-auto">
          <!-- block1 -->
          <div class="block1 hov-img-zoom pos-relative m-b-30">
            <img src="https://baamboo.com/wp-content/uploads/2020/06/nuoc-hoa-hong-thayers-7.jpg" alt="IMG-BENNER" />

            <div @click="summitCategory(4)" class="block1-wrapbtn w-size2">
              <!-- Button -->
              <router-link to="/products"  class="flex-c-m size2 m-text2 bg3 hov1 trans-0-4">
                TONER
              </router-link>
            </div>
          </div>

          <!-- block1 -->
          <div class="block1 hov-img-zoom pos-relative m-b-30">
            <img src="https://cf.shopee.vn/file/9eff84beecee212da4c9433e5cca3004" alt="IMG-BENNER" />

            <div  @click="summitCategory(2)" class="block1-wrapbtn w-size2">
              <router-link to="/products"  class="flex-c-m size2 m-text2 bg3 hov1 trans-0-4">
                MASK
              </router-link>
            </div>
          </div>
        </div>

        <div class="col-sm-10 col-md-8 col-lg-4 m-l-r-auto">
          <!-- block1 -->
          <div class="block1 hov-img-zoom pos-relative m-b-30">
            <img src="https://tudienlamdep.org/wp-content/uploads/2019/09/kem-chong-nang-nuskin-sunright-spf50-3.jpg" alt="IMG-BENNER" />

            <div  @click="summitCategory(3)" class="block1-wrapbtn w-size2">
              <!-- Button -->
              <router-link to="/products"  class="flex-c-m size2 m-text2 bg3 hov1 trans-0-4">
                FACE SCREAM
              </router-link>
            </div>
          </div>

          <!-- block2 -->
          <div class="block2 wrap-pic-w pos-relative m-b-30">
            <img src="https://classbento.com.au/images/class/natural-skincare-making-private-class-600.jpg?1589939169" alt="IMG" />

            <div class="block2-content sizefull ab-t-l flex-col-c-m">
              <h4 class="m-text4 t-center w-size3 p-b-8">
                Sign up & get 20% off
              </h4>

              <p style="color: black" class="t-center w-size4">
                Be the frist to know about the latest fashion news and get
                exclu-sive offers
              </p>

              <div class="w-size2 p-t-25">
                <router-link
                    to="/login"
                  class="flex-c-m size2 bg4 bo-rad-23 hov1 m-text3 trans-0-4"
                >
                  Sign Up
                </router-link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>


export default {
  name: "SectionCategory",
  methods: {
    summitCategory (category) {
          this.$store.dispatch("products/getProductByCategory", category)
    }
  }
};
</script>