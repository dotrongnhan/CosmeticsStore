import axios from "axios";

const state = () => ({
  products: [],
  isLoading: false,
  addToCartResult: "",
  totalItems: 0,
  isShowCartDropdown: false,
});

const getters = {
  totalItems(state) {
    return state.products.reduce(
      (total, product) => total + product.quantity,
      0
    );
  },

  subTotal(state) {
    return state.products.reduce(
      (totalPrice, product) => totalPrice + product.quantity * product.product.price,
      0
    );
  },
};

const actions = {
  async getOrderItemsByUserId({ commit}, id) {
    try {
      const res = await axios.get(`orders/${id}`, {withCredentials: true});
      commit("SET_PRODUCTS", res.data)
    } catch (e) {
      console.log(e);
    }

  },
  async createOrderItem(_, props) {
    console.log(props)
    try {
      const res = await axios.post('orders', props, {withCredentials: true});
      console.log("register success" + res)
      // commit("setRegisterSuccess", true);
      // commit("setRegisterMessage", "");
    } catch (e) {
      console.log(e)
      // commit("setRegisterMessage", e.message === "Request failed with status code 400" ? "Email already in use" : "Create new account is failed");
      // commit("setRegisterSuccess", false);
    }
  },
};

const mutations = {
  SET_PRODUCTS(state, products) {
    state.products = products;
  },
  setShowCartDropdown(state, status) {
    state.isShowCartDropdown = status;
  },

  updateProductQuantity(state, { productId, value }) {
    const product = state.products.find((p) => p.id === productId);
    const index = state.products.findIndex(item => item.id === productId)
    value = Number(value);

    if (value > 1) {
      product.quantity = value;
    } else {
      state.products.splice(index,1)
    }

    product.totalPrice = product.price * product.quantity;
  },
  addProductToCart(state, {product, quantity}) {
    const isExist = state.products.findIndex(item => item.product.id === product.id)
    if(quantity === 0 ) return
    if (isExist === -1) {
      state.products.push({product: product, quantity: quantity});
    } else {
      state.products[isExist].quantity = state.products[isExist].quantity + quantity
    }
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
