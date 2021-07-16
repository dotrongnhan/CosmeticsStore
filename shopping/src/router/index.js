import { createRouter, createWebHistory } from "vue-router";

const routes = [
  { path: "/", component: () => import("../pages/home/HomePage") },
  { path: "/products", component: () => import("../pages/products/ProductsPage") },
  { path: "/products/:id", component: () => import("../pages/product-detail/ProductDetailPage") },
  { path: "/about", component: () => import("../pages/about/AboutPage") },
  { path: "/contact", component: () => import("../pages/contact/ContactPage") },
  { path: "/cart", component: () => import("../pages/cart/CartPage") },
  { path: "/login", component: () => import("../pages/login/LoginPage") },
  { path: "/register", component: () => import("../pages/register/RegisterPage") },
  { path: "/payment", component: () => import("../pages/payment/Payment")},
  { path: "/confirm-checkout", component: () => import("../pages/confirm-checkout/ConfirmCheckout")},
  {
    path: "/user",
    component: () => import("../pages/user/UserPage"),
    children: [
      { path: "", component: () => import("../pages/user/ProfilePage") },
      { path: "profile", component: () => import("../pages/user/ProfilePage") },
      { path: "settings", component: () => import("../pages/user/SettingsPage") },
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
