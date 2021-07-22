import axios from "axios";

const state = () => ({
  products: [],
  count: 0,
  product: {},
  search: "",
  category: { value: 'Default', label: 'Default category' },
  sort: 0,
});

const getters = {

};

const actions = {
  getProducts: async ({commit}, {category = 'Default', sortType = 'Default', offset = 1}) => {
    try {
      const res = await axios.get(`products?sortType=${sortType}&offset=${offset}&category=${category}`, )
      commit("GET_PRODUCTS", res.data)
    } catch (e) {
      console.log(e)
    }
  },
  getProductById: async ({commit}, id) => {
    try {
      const res = await axios.get(`products/${id}`)
      commit("GET_PRODUCTS_BY_ID", res.data)
    } catch (e) {
      console.log(e)
    }
  },
  deleteProduct: async ({commit}, id) => {
    try {
      await axios.delete(`products/${id}`, {withCredentials: true})
      await commit('DELETE_PRODUCT', id)
    } catch (e) {
      console.log(e)
    }
  },
  createProduct: async ( _, product) => {
    try {
      await axios.post("products", product, {withCredentials: true})
    } catch (e) {
      console.log(e)
    }
  },
  updateProduct: async (_, payload) => {
    try {
      await axios.post(`products/${payload.id}`,payload.product, {withCredentials: true})
    } catch (e) {
      console.log(e)
    }
  }
};
const mutations = {
  DELETE_PRODUCT (state, id) {
    const index = state.products.findIndex(product => product.id === id)
    state.products.splice(index, 1)
  },
  GET_PRODUCTS (state, data) {
    state.products = data.Products
    state.count = data.Count
  },
  GET_PRODUCT_BY_CATEGORY (state, category) {
    state.category = category
    },
  GET_PRODUCTS_BY_ID: (state, product) => {
    state.product = product
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
