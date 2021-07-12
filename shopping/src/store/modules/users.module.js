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
  async login({ commit }, user) {
    try {
      const res = await axios.post("login", user)
      const User = res.data.User
      console.log( User)
      // document.cookie = "Bearer=" + res.data.Token;
      localStorage.setItem("User", JSON.stringify(User) )
      commit("setLoginSuccess", true);
      commit("setLoginMessage", "");
      commit("setUser", User );
    } catch (error) {
      console.log(error)
      commit("setLoginSuccess", false);
      commit(
          "setLoginMessage", "User name or password is wrong!"
      );
    }
  },

  async logout({ commit}){
    await axios.post("logout");
    await commit("logout")
    
},

  async register({ commit }, user) {
    try {
      const res = await axios.post('http://127.0.0.1:3000/api/user/register', user)
      console.log(res)
      commit("setRegisterSuccess", true);
      commit("setRegisterMessage", "");
    } catch (e) {
      console.log(e)
      commit("setRegisterMessage", e.message === "Request failed with status code 400" ? "Email already in use" : "");
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
