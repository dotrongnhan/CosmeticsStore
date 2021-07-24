<template>
  <div class="header-wrapicon2" @mouseover="toggleCartDropdown" @mouseleave="toggleCartUp">
    <img
      src="@/assets/images/icons/icon-header-02.png"
      class="header-icon1 js-show-header-dropdown"
      alt="ICON"
    />
    <span class="header-icons-noti">{{ totalItems }}</span>

    <div
      class="header-cart header-dropdown"
      :class="{ 'show-header-dropdown': isShowCartDropdown && isLoginSuccess }"
    >
      <p>{{ addToCartResult }}</p>

      <ul class="header-cart-wrapitem">
        <li
          class="header-cart-item"
          v-for="(product, index) in products"
          :key="index"
        >
          <div class="header-cart-item-img">
            <img :src="product.product.image" alt="IMG" />
          </div>

          <div class="header-cart-item-txt">
            <a href="#" class="header-cart-item-name">
              {{ product.product.product_name }}
            </a>

            <span class="header-cart-item-info">
              {{ product.quantity }} x {{ currency(product.product.price) }}
            </span>
          </div>
        </li>
      </ul>

      <div class="header-cart-total">Total: {{ currency(subTotal) }}</div>

      <div class="header-cart-buttons">
        <div class="header-cart-wrapbtn">
          <router-link
            to="/cart"
            class="flex-c-m size1 bg1 bo-rad-20 hov1 s-text1 trans-0-4"
          >
            View Cart
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapGetters } from "vuex";
import { currency } from "@/utils/currency";

export default {
  name: "HeaderCartDropdown",

  computed: {
    ...mapState("cart", ["products", "isShowCartDropdown", "addToCartResult"]),
    ...mapGetters("cart", ["totalItems", "subTotal"]),
    ...mapState("users",["isLoginSuccess"])
  },

  methods: {
    toggleCartDropdown() {
      this.$store.commit("cart/setShowCartDropdown", true);
    },
    toggleCartUp() {
      this.$store.commit("cart/setShowCartDropdown", false);
    },
    currency,
  },
};
</script>
