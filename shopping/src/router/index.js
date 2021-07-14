import { createRouter, createWebHistory } from "vue-router";

const routes = [
  { path: "/", component: () => require("../pages/home/HomePage") },
  { path: "/products", component: () => require("../pages/products/ProductsPage") },
  { path: "/products/:id", component: () => require("../pages/product-detail/ProductDetailPage") },
  { path: "/about", component: () => require("../pages/about/AboutPage") },
  { path: "/contact", component: () => require("../pages/contact/ContactPage") },
  { path: "/cart", component: () => require("../pages/cart/CartPage") },
  { path: "/login", component: () => require("../pages/login/LoginPage") },
  { path: "/register", component: () => require("../pages/register/RegisterPage") },
  {path: "/payment", component: () => require("../pages/payment/Payment")},
  {path: "/confirm-checkout", component: () => require("../pages/confirm-checkout/ConfirmCheckout")},
  {
    path: "/user",
    component: () => require("../pages/user/UserPage"),
    children: [
      { path: "", component: () => require("../pages/user/ProfilePage") },
      { path: "profile", component: () => require("../pages/user/ProfilePage") },
      { path: "settings", component: () => require("../pages/user/SettingsPage") },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  scrollBehavior() {
    return { top: 0 };
  },
  routes,
});

export default router;
