import { createStore } from "vuex";
import users from "./modules/users.module";
import products from "./modules/products.module";
import cart from "./modules/cart.module";


const store = createStore({
  modules: {
    users,
    products,
    cart
  },
});

export default store;
