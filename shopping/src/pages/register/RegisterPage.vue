<template>
  <div>
    <!-- Title Page -->
    <section
      class="bg-title-page p-t-40 p-b-50 flex-col-c-m"
      :style="{
        backgroundImage:
          'url(' + require('@/assets/background-register.jpeg') + ')',
      }"
    >
      <h2 class="l-text2 t-center">Register account</h2>
    </section>

    <!-- content page -->
    <section class="bgwhite p-t-66 p-b-60">
      <div class="container">
        <div class="register-form m-auto">
          <template v-if="!isRegisterSuccess">
            <Form @submit="onSubmit" :validation-schema="schema">
              <p class="m-b-10 text-center text-danger">
                {{ registerMessage }}
              </p>

              <div class="bo4 of-hidden size15 m-b-10">
                <Field
                    v-model="user.full_name"
                  name="full_name"
                  type="text"
                  placeholder="Your Name"
                  class="sizefull s-text7 p-l-22 p-r-22"
                  :disabled="isLoading"
                />
              </div>

              <ErrorMessage
                name="full_name"
                class="text-danger m-b-20 d-block"
              />

              <div class="bo4 of-hidden size15 m-b-10">
                <Field
                    v-model="user.email"
                  name="email"
                  type="text"
                  placeholder="Email"
                  class="sizefull s-text7 p-l-22 p-r-22"
                  :disabled="isLoading"
                />
              </div>

              <ErrorMessage name="email" class="text-danger m-b-20 d-block" />

              <div class="bo4 of-hidden size15 m-b-10">
                <Field
                    v-model="user.password"
                  name="password"
                  type="password"
                  placeholder="Password"
                  class="sizefull s-text7 p-l-22 p-r-22"
                  :disabled="isLoading"
                />
              </div>

              <ErrorMessage name="password" class="text-danger m-b-20 d-block" />

              <div class="bo4 of-hidden size15 m-b-10">
                <Field
                  name="confirm_password"
                  type="password"
                  placeholder="Confirm Password"
                  class="sizefull s-text7 p-l-22 p-r-22"
                  :disabled="isLoading"
                />
              </div>

              <ErrorMessage
                name="confirm_password"
                class="text-danger m-b-20 d-block"
              />

              <div class="w-size25 m-auto">
                <button
                    class="flex-c-m size2 bg1 bo-rad-23 hov1 m-text3 trans-0-4 m-t-20"
                  :class="{ disabled: isLoading }"
                  :disabled="isLoading"
                >
                  <span v-show="isLoading" data-loader="ball-scale"></span>
                  Register
                </button>
              </div>
            </Form>

            <br />
            <hr />

            <p class="text-center">
              <router-link to="/login"
                >Have an account? Login now!
              </router-link>
            </p>
          </template>

          <template v-else>
            <p class="text-center">{{ message }}</p>

            <router-link to="/login"
              ><h4 class="text-center">Go to Login Page »</h4></router-link
            >
          </template>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import { mapState } from "vuex";
import {Form, Field, ErrorMessage} from "vee-validate";
import * as yup from "yup";
import * as Yup from "yup";

export default {
  name: "RegisterPage",

  components: {
    Form,
    Field,
    ErrorMessage,
  },

  data() {
    const schema = yup.object().shape({
      full_name: yup
        .string()
        .required("Your name is required!"),
      email: yup
        .string()
        .required("Email is required!")
        .email("Email is invalid!")
        .max(50, "Email must be maximum 50 characters!"),
      password: yup
        .string()
        .required("Password is required!")
        .min(6, "Password must be at least 6 characters!")
        .max(40, "Password must be maximum 40 characters!"),
      confirm_password: yup
      .string()
      .oneOf([Yup.ref('password'), null], 'Passwords must match')
    });

    return {
      isLoading: false,
      message: "",
      schema,
      user: {
        full_name: '',
        email: '',
        password: ''
      }
    };

  },

  computed: mapState("users", [
    "isLoginSuccess",
    "isRegisterSuccess",
    "registerMessage",
  ]),

  created() {
    if (this.isLoginSuccess) {
      this.$router.replace("/user");
    }
  },

  methods: {
    onSubmit() {
      this.$store.dispatch("users/register", {
        full_name: this.user.full_name,
        email: this.user.email,
        password: this.user.password,
      });

      this.isLoading = false;

      if (this.isRegisterSuccess) {
        this.message = "Register successfully.";
      }
    },
  },
};
</script>

<style scoped>
.register-form {
  max-width: 500px;
}
button.disabled {
  opacity: 0.7;
  cursor: not-allowed;
}
button.disabled:hover {
  background-color: #222222;
}
</style>