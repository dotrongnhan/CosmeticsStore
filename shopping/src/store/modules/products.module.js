import axios from "axios";

const state = () => ({
  products: [],
  count: 0,
  product: {},
  search: "",
  category: { value: 'default', label: 'Default category' },
  sort: 0,
});

const getters = {
  currentCategory: state => {
    return state.category
  }
};

const actions = {
  getProducts: async ({commit}, {category, sortType, offset = 1}) => {
    console.log(category, sortType, offset)
    try {
      const res = await axios.get(`products?sortType=${sortType}&offset=${offset}&category=${category}`, )
      commit("GET_PRODUCTS", res.data)
    } catch (e) {
      console.log(e)
    }
  },
  getProductById: async ({commit}, id) => {
    console.log(id)
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
    console.log(category)
    state.category = category
    },
  GET_PRODUCTS_BY_ID: (state, product) => {
    state.product = product
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
