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
  async updateCart({commit}, {userId, product, quantity, replace = "", isPaid = 0}) {
      try {
        const res = await axios.post("order/upsert", {user_id: userId, product_id: product.id, quantity: quantity, replace: replace, is_paid: isPaid})
        console.log("ok", res)
        if (replace === "") {
          commit("addProductToCart", {product: product, quantity: quantity})
        }
      } catch (e) {
        console.log(e)
      }
  },
  async getCartById({commit}, id) {
    try {
      const res = await axios.get(`/orders/user/${id}`)
      commit("CALL_CART_FROM_SV", res.data)
      console.log(res)
    } catch (e) {
      console.log(e)
    }
  },
  async deleteOrderItem({commit}, id) {
    try {
      const res = axios.delete(`/orders/${id}`)
      console.log(res)
      commit("REMOVE_ORDER", id)
    } catch (e) {
      console.log(e)
    }
  }

};

const mutations = {
  CALL_CART_FROM_SV(state, cart) {
      state.products = cart
  },
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
  REMOVE_ORDER(state, id) {
    state.products = state.products.filter(product => product.id !== id)
  },
  REMOVE_CART(state) {
    state.products = []
    state.totalItems = 0
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
