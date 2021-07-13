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
};

const mutations = {
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
