
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
  async getUserExits({commit}) {
    try {
      const res = await axios.get("user", {withCredentials: true})
      commit("setLoginSuccess", true);
      commit("setUser", res.data );
    } catch (e) {
      console.log(e)
    }
  }
  ,
  async login({ commit }, user) {
    try {
      const res = await axios.post("login", user, {withCredentials: true})
      commit("setLoginSuccess", true);
      commit("setLoginMessage", "");
      commit("setUser", res.data.User );
    } catch (error) {
      console.log(error)
      commit("setLoginSuccess", false);
      commit(
          "setLoginMessage", "User name or password is wrong!"
      );
    }
  },
  async logout({ commit}){
    await axios.post("logout","hello", {withCredentials: true});
    await commit("logout")
    
},

  async register({ commit }, user) {
    try {
      const res = await axios.post('register', user)
      console.log("register success" + res)
      commit("setRegisterSuccess", true);
      commit("setRegisterMessage", "");
    } catch (e) {
      console.log(e)
      commit("setRegisterMessage", e.message === "Request failed with status code 400" ? "Email already in use" : "Create new account is failed");
      commit("setRegisterSuccess", false);
    }
  },
};

const mutations = {
  setUser(state, user) {
    state.user = user;
  },

  setLoginSuccess(state, status) {
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
    state.isLoginSuccess = false;
    state.isRegisterSuccess = false;
    localStorage.removeItem("User");
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};