
import axios from "axios";

const state = () => ({
  users: [],
  user: {},
  isLoginSuccess: false,
  loginMessage: "",
  isRegisterSuccess: false,
  isUpdateSuccess: false,
  isChangePassSuccess: false,
  registerMessage: "",
  updateMessage: "",
  changePassMessage: "",
  isShowUserDropdown: false,
  isDisplayForm: false,
});

const getters = {};

const actions = {
  getUsers: async ({commit}) => {
    try {
      const res = await axios.get('users', {withCredentials: true});
      console.log(res);
      commit("GET_USERS", res.data)
      commit("SET_USERS", res.data)
    } catch (e) {
      console.log(e)
    }
  },
  deleteUser: async ({commit}, id) => {
    let decision = confirm("Are you sure you want to delete?");
      if (decision === true) {
    try {
      await axios.delete(`users/${id}`, {withCredentials: true})
      await commit('DELETE_USER', id)
    } catch (e) {
      console.log(e)
    }
  }
  },
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

  async updateProfile({commit}, user) {
    try {
        const res = await axios.put('user', user, {withCredentials: true});
        console.log(res);
        commit("setUpdateSuccess", true);
        commit("setUpdateMessage", "");
    } catch (e) {
      console.log(e)
      commit("setUpdateMessage", e.message === "Request failed with status code 400" ? "????" : "????");
      commit("setUpdateSuccess", false);
    }
  },
  async changePassword({commit}, user) {
    try {
        const res = await axios.put('user/password', user, {withCredentials: true});
        console.log(res);
        commit("setChangePassSuccess", true);
        commit("setChangePassMessage", "Change password successfully!");
    } catch (e) {
      // console.log(e)
      commit("setChangePassMessage", e.message === "Request failed with status code 400" ? "????" : "????");
      commit("setChangePassSuccess", false);
    }
  }
};

const mutations = {
  GET_USERS (state, data) {
    state.users = data.Users
  },
  DELETE_USER(state, id) {
    const index = state.users.findIndex(user => user.id === id)
    state.users.splice(index, 1)    
  },
  setShowUserDropdown(state, status) {
    state.isShowUserDropdown = status;
  },
  setUser(state, user) {
    state.user = user;
  },
  SET_USERS(state, users) {
    state.users = users;
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

  setUpdateSuccess(state, status) {
    state.isUpdateSuccess = status;
  },

  setUpdateMessage(state, message) {
    state.updateMessage = message;
  },

  setChangePassSuccess(state, status) {
    state.isChangePassSuccess = status;
  },

  setChangePassMessage(state, message) {
    state.changePassMessage = message;
  },

  logout(state) {
    state.user = {};
    state.isLoginSuccess = false;
    state.isRegisterSuccess = false;
    localStorage.removeItem("User");
  },
  displayForm(state) {
   if(state.isDisplayForm == false) {
    state.isDisplayForm = true;
   }else {
     state.isDisplayForm = false;
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