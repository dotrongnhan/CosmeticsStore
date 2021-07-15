<template>
  <header class="header1">
    <div class="container-menu-header">
      <div class="wrap_header">
        <router-link to="/" class="logo">
          <img src="https://media.istockphoto.com/vectors/women-beauty-icon-vector-id877241778?b=1&k=6&m=877241778&s=612x612&w=0&h=kK7KTz70gcdMxzD3rUjqKagoTdw6VbS-p5EgFar_BME=" alt="IMG-LOGO" />
        </router-link>

        <div class="wrap_menu">
          <nav class="menu">
            <ul class="main_menu">
              <MenuItem to="/" label="Home" />

              <MenuItem to="/products" label="Products" />

              <MenuItem to="/about" label="About" />

              <MenuItem to="/contact" label="Contact" />
            </ul>
          </nav>
        </div>

        <div class="header-icons">
          <!-- <router-link :to="isLoginSuccess ? '/user' : '/login'">
            <img
              :src="user?.avatar ? user.avatar : defaultAvatar"
              class="header-icon1"
              style="width: 27px; border-radius: 50%"
              alt="Avatar"
            />
          </router-link> -->
          <HeaderUserDropdown />

          <span class="linedivide1"></span>

          <HeaderCartDropdown />
        </div>
      </div>
    </div>

    <!-- Header Mobile -->
    <div class="wrap_header_mobile">
      <!-- Logo moblie -->
      <router-link to="/" class="logo-mobile">
        <img src="https://media.istockphoto.com/vectors/women-beauty-icon-vector-id877241778?b=1&k=6&m=877241778&s=612x612&w=0&h=kK7KTz70gcdMxzD3rUjqKagoTdw6VbS-p5EgFar_BME=" alt="IMG-LOGO" />
      </router-link>

      <!-- Button show menu -->
      <div class="btn-show-menu">
        <!-- Header Icon mobile -->
        <div class="header-icons-mobile">
          <a href="#" class="header-wrapicon1 dis-block">
            <img
              :src="user?.avatar ? user.avatar : defaultAvatar"
              class="header-icon1"
              alt="Avatar"
            />
          </a>

          <span class="linedivide2"></span>

          <HeaderCartDropdown />
        </div>

        <div
          class="btn-show-menu-mobile hamburger hamburger--squeeze"
          :class="{ 'is-active': isShowMenuMobileDropdown }"
          @click="toggleMenuMobileDropdown"
        >
          <span class="hamburger-box">
            <span class="hamburger-inner"></span>
          </span>
        </div>
      </div>
    </div>

    <!-- Menu Mobile -->
    <div
      class="wrap-side-menu"
      :style="{ display: isShowMenuMobileDropdown ? 'block' : 'none' }"
    >
      <nav class="side-menu">
        <ul class="main-menu">
          <li class="item-topbar-mobile p-l-10">
            <div class="topbar-social-mobile">
              <a href="#" class="topbar-social-item fa fa-facebook"></a>
              <a href="#" class="topbar-social-item fa fa-instagram"></a>
              <a href="#" class="topbar-social-item fa fa-pinterest-p"></a>
              <a href="#" class="topbar-social-item fa fa-snapchat-ghost"></a>
              <a href="#" class="topbar-social-item fa fa-youtube-play"></a>
            </div>
          </li>

          <li class="item-menu-mobile" @click="closeMenuMobileDropdown">
            <router-link to="/">Home</router-link>
          </li>

          <li class="item-menu-mobile" @click="closeMenuMobileDropdown">
            <router-link to="/products">Products</router-link>
          </li>

          <li class="item-menu-mobile" @click="closeMenuMobileDropdown">
            <router-link to="/about">About</router-link>
          </li>

          <li class="item-menu-mobile" @click="closeMenuMobileDropdown">
            <router-link to="/contact">Contact</router-link>
          </li>
        </ul>
      </nav>
    </div>
  </header>
</template>

<script>
import {mapActions, mapMutations, mapState} from "vuex";
import defaultAvatar from "@/assets/images/icons/icon-header-01.png";
import MenuItem from "../components/MenuItem.vue";
import HeaderCartDropdown from "../components/HeaderCartDropdown.vue";
import HeaderUserDropdown from "../components/HeaderUserDropdown.vue";
import {getCookie} from "../utils/getCookieByName";

export default {
  name: "MainHeader",

  components: {
    MenuItem,
    HeaderCartDropdown,
    HeaderUserDropdown
  },

  data() {
    return {
      defaultAvatar,
      language: {
        selected: {},
      },
      isShowMenuMobileDropdown: false,
    };
  },
  created() {
    this.getProducts({})
    if (getCookie("jwt")) {
      this.getUserExits()
    }
  },

  computed: mapState("users", ["isLoginSuccess", "user"]),

  methods: {
    getCookie,
    ...mapActions("users",["getUserExits"]),
    ...mapActions("products", ["getProducts"]),
    ...mapMutations("users", ["updateUser"]),
    toggleMenuMobileDropdown() {
      this.isShowMenuMobileDropdown = !this.isShowMenuMobileDropdown;
    },

    closeMenuMobileDropdown() {
      this.isShowMenuMobileDropdown = false;
    },
  },
};
</script>

