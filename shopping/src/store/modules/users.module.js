import axios from "axios";

const state = () => ({
  user: {},
  isLoginSuccess: false,
  loginMessage: "",
  isRegisterSuccess: false,
  registerMessage: "",
});

const getters = {};

const actions = {
  async updateUser() {
    try {
      const res = await axios.get("/user")
      console.log(res)
    }catch (e) {
      console.log(e)
    }
  },
  async login({ commit }, user) {
    try {
      const res = await axios.post("login", user, {withCredentials: true})
      const User = res.data.User
      commit("setLoginSuccess", true);
      commit("setLoginMessage", "");
      commit("setUser", User );
    } catch (error) {
      console.log(error)
      commit("setLoginSuccess", false);
      commit("setLoginMessage", "User name or password is wrong!");
    }
  },

  async register({ commit }, user) {
    console.log(user)
    try {
      const res = await axios.post('register', user)
      console.log(res)
      console.log(1)
      commit("setRegisterSuccess", true);
      commit("setRegisterMessage", "");
    } catch (e) {
      commit("setRegisterMessage", e.message === "Request failed with status code 405" ? "Email already in use" : "");
      commit("setRegisterSuccess", false);
    }
  },
};

const mutations = {
  updateUser(state, user) {
    state.user = user
    state.isLoginSuccess = true
  },
  setUser(state, user) {
    state.user = user;
  },

  setLoginSuccess(state, status) {
    console.log(status)
    state.isLoginSuccess = status;
  },

  setLoginMessage(state, message) {
    state.loginMessage = message;
  },

  setRegisterSuccess(state, status) {
    state.isRegisterSuccess = status;
  },

  setRegisterMessage(state, message) {
    state.registerMessage = message;
  },

  logout(state) {
    state.user = {};
    localStorage.removeItem("User")
    state.isLoginSuccess = false;
    state.isRegisterSuccess = false;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
