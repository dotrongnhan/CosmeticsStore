import axios from "axios";

const state = () => ({
  products: [],
  count: 0,
  product: {},
  search: "",
  category: { value: '', label: 'Default category' },
  sort: 0,
});

const getters = {
  currentCategory: state => {
    return state.category
  }
};

const actions = {
  getProducts: async ({commit}, {sortType, prop, offset}) => {
    console.log(sortType, prop, offset)
    try {
      const res = await axios.get(`products?sortType=${sortType}&prop=${prop}&offset=${offset}`, )
      commit("GET_PRODUCTS", res.data)
    } catch (e) {
      console.log(e)
    }
  },
  getProductById: async ({commit}, id) => {
    try {
      const res = await axios.get(`products/${id}`)
      console.log(res.data)
      commit("GET_PRODUCTS_BY_ID", res.data)
    } catch (e) {
      console.log(e)
    }
  },
  getProductByCategory: async ({commit}, category) => {
    const res = await axios.get(`http://127.0.0.1:3000/api/category/${category}`)
    commit('GET_PRODUCT_BY_CATEGORY', res.data)
  }
};
const mutations = {
  GET_PRODUCTS (state, data) {
    state.products = data.Products
    state.count = data.Count
  },
  GET_PRODUCT_BY_CATEGORY (state, category) {
    const categoryCurrent = category.category_name
    state.products = category.products
    if(categoryCurrent === 'mask') {
      state.category = { value: 2, label: 'Mask' }
      console.log(state.category)
    }
    if(categoryCurrent === 'face scream') {
      state.category = { value: 3, label: 'Face Scream' }
    }
    if(categoryCurrent === 'toner') {
      state.category = { value: 4, label: 'Toner' }
    }
    if(categoryCurrent === 'lipstick') {
      state.category = { value: 5, label: 'Lipstick' }
    }
    if(categoryCurrent === 'serum') {
      state.category = { value: 6, label: 'Serum' }
    }
    },
  GET_PRODUCTS_BY_ID: (state, product) => {
    state.product = product
  },
  SORT_PRODUCT:(state, direction) => {
    if(direction === -1) {
      state.products = state.products.sort((a,b) => a.price - b.price)
    }
    if(direction === 1) {
      state.products = state.products.sort((a,b) => b.price - a.price)
    }
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
