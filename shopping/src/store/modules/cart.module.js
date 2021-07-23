import axios from "axios";

const state = () => ({
  allProducts: [],
  count: 0,
  customerProducts: [],
  products: [],
  isLoading: false,
  addToCartResult: "",
  isDisplayFormOrder: false,
  isShowCartDropdown: false,
});

const getters = {
  totalItems(state) {
    if(state.products === null || state.products === undefined ) return 0
    return state.products.reduce(
      (total, product) => total + product.quantity,
      0
    );
  },

  subTotal(state) {
    if(state.products === null || state.products === undefined ) return 0
    return state.products.reduce(
      (totalPrice, product) => totalPrice + product.quantity * product.product.price,
      0
    );
  },

  subTotal1(state) {
    if(state.customerProducts === null || state.customerProducts === undefined ) return 0
    return state.customerProducts.reduce(
      (totalPrice, customerProduct) => totalPrice + customerProduct.quantity * customerProduct.product.price,
      0
    );
  },
  getProductsInCart(state) {
      return state.products.map(item => item.product.id)
  }
};

const actions = {
  changeStatusOrderItems: async (_, {products, userId}) => {
      await axios.post('orders/change', {user_id: userId, products: products })
  },
  getAllOrders: async ({commit}, {offset = 1}) => {
    try {
      const res = await axios.get(`orders?offset=${offset}`, {withCredentials: true});
      console.log(res);
      commit("SET_ORDERS", res.data)
    } catch (e) {
      console.log(e)
    }
  },
  deleteOrder: async ({commit}, id) => {
    let decision = confirm("Are you sure you want to delete?");
      if (decision === true) {
    try {
      await axios.delete(`orders/admin/${id}`, {withCredentials: true})
      await commit('DELETE_ORDER', id)
    } catch (e) {
      console.log(e)
    }
  }
  },
  async updateCart({commit}, {userId, product, quantity, replace = "", isPaid = 0}) {
      try {
        const res = await axios.post("order/upsert", {user_id: userId, product_id: product.id, quantity: quantity, replace: replace, is_paid: isPaid})
        console.log("ok", res)
        if (replace === "") {
          commit("addProductToCart", {userId: userId, productId: product.id, product: product, quantity: quantity})
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
  async getCartByCustomerId({commit}, id) {
    try {
      const res = await axios.get(`/orders/user/${id}`)
      commit("CALL_CCART_FROM_SV", res.data)
      console.log(res)
    } catch (e) {
      console.log(e)
    }
  },
  async deleteOrderItem({commit}, product) {
    try {
      const res = await axios.delete(`/orders/${product.product.id}`, {withCredentials: true})
      console.log(res)
      commit("REMOVE_ORDER", product.product.id)
    } catch (e) {
      console.log(e)
    }
  }

};

const mutations = {
  SET_ORDERS (state, data) {
    state.allProducts = data.OrderItems;
    state.count = data.Count;
  },
  DELETE_ORDER(state, id) {
    const index = state.allProducts.findIndex(allProduct => allProduct.id === id)
    state.allProducts.splice(index, 1)    
  },
  CALL_CART_FROM_SV(state, cart) {
      state.products = cart
  },
  CALL_CCART_FROM_SV(state, cart) {
    state.customerProducts = cart
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
  addProductToCart(state, {userId, productId, product, quantity}) {
    console.log(quantity)
    let isExist = state.products?.findIndex(item => item.product.id === product.id)
    if (isExist === undefined) {
      console.log("isExist === undefined")
      isExist = -1
      state.products = []
    }
    if(quantity === 0 ) return
    if (isExist === -1) {
      state.products.push({user_id: userId, product_id: productId, product: product, quantity: quantity});
    } else {
      console.log(isExist)
      isExist = state.products.findIndex(item => item.product.id === product.id)
      state.products[isExist].quantity = state.products[isExist].quantity + quantity
    }
  },
  displayForm(state) {
    if(state.isDisplayFormOrder == false) {
     state.isDisplayFormOrder = true;
    }else {
      state.isDisplayFormOrder = false;
    }
   },
  REMOVE_ORDER(state, id) {
    state.products = state.products.filter(product => product.product.id !== id)
  },
  REMOVE_CART(state) {
    state.products = []
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
