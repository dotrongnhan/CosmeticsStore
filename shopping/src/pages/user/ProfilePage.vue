<template>
  <div class="container emp-profile">
    <form method="post">
      <div class="row">
        <div class="col-md-4 pb-2">
          <div class="profile-img">
            <img :src="user.avatar" alt=""/>
          </div>
        </div>
        <div class="col-md-6">
          <div class="profile-head">
            <h3>
              {{user.full_name}}
            </h3>
            <p class="proile-rating">Position : <span>{{user.role_id === 1 ? "Admin" : "Customer"}}</span></p>
            <ul class="nav nav-tabs" id="myTab" role="tablist">
              <li class="nav-item">
                <a class="nav-link active" id="home-tab" data-toggle="tab" href="#home" role="tab" aria-controls="home" aria-selected="true">About</a>
              </li>
            </ul>
          </div>
        </div>
        <div class="col-md-2">
        </div>
      </div>
      <div class="row" id="profile">
        <div class="col-md-8">
          <div class="tab-content profile-tab" id="myTabContent">
            <!-- <div class="tab-pane fade show active" id="home" role="tabpanel" aria-labelledby="home-tab">
              <div class="row">
                <div class="col-md-6">
                  <label>Full name</label>
                </div>
                <div class="col-md-6">
                  <p>{{user.full_name}}</p>
                </div>
              </div>
              <div class="row">
                <div class="col-md-6">
                  <label>Gender</label>
                </div>
                <div class="col-md-6">
                  <p>{{user.gender}}</p>
                </div>
              </div>
              <div class="row">
                <div class="col-md-6">
                  <label>Date of Birth</label>
                </div>
                <div class="col-md-6">
                  <p>{{user.date_of_birth}}</p>
                </div>
              </div>
              <div class="row">
                <div class="col-md-6">
                  <label>Email</label>
                </div>
                <div class="col-md-6">
                  <p>{{user.email}}</p>
                </div>
              </div>
              <div class="row">
                <div class="col-md-6">
                  <label>Address</label>
                </div>
                <div class="col-md-6">
                  <p>{{user.address}}</p>
                </div>
              </div>
              <div class="row">
                <div class="col-md-6">
                  <label>Phone</label>
                </div>
                <div class="col-md-6">
                  <p>{{user.phone}}</p>
                </div>
              </div>
            </div> -->

            <!-- <div class="tab-pane fade" id="profile" role="tabpanel" aria-labelledby="profile-tab">
              <div class="row">
                <div class="col-md-6">
                  <label>Experience</label>
                </div>
                <div class="col-md-6">
                  <p>Expert</p>
                </div>
              </div>
              <div class="row">
                <div class="col-md-6">
                  <label>Hourly Rate</label>
                </div>
                <div class="col-md-6">
                  <p>10$/hr</p>
                </div>
              </div>
              <div class="row">
                <div class="col-md-6">
                  <label>Total Projects</label>
                </div>
                <div class="col-md-6">
                  <p>230</p>
                </div>
              </div>
              <div class="row">
                <div class="col-md-6">
                  <label>English Level</label>
                </div>
                <div class="col-md-6">
                  <p>Expert</p>
                </div>
              </div>
              <div class="row">
                <div class="col-md-6">
                  <label>Availability</label>
                </div>
                <div class="col-md-6">
                  <p>6 months</p>
                </div>
              </div>
              <div class="row">
                <div class="col-md-12">
                  <label>Your Bio</label><br/>
                  <p>Your detail description</p>
                </div>
              </div>
            </div> -->
          
            <Form @submit="onSubmit" :validation-schema="schema" style="margin-left: 390px; margin-top: -100px; width : 74%;">
                <div>
                <button
                    style = "width: 30%; margin-left: 390px"
                    @click="onSubmit"
                    class="flex-c-m size2 bg1 bo-rad-23 hov1 m-text3 trans-0-4 m-t-20"
                  :class="{ disabled: isLoading }"
                  :disabled="isLoading"
                >
                  <span v-show="isLoading" data-loader="ball-scale"></span>
                  Edit Profile
                </button>
              </div>
                  <div class="col-md-6">
                    <label>Full name</label>
                  </div>
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
                  <div class="col-md-6">
                    <label>Email</label>
                  </div>
              <div class="bo4 of-hidden size15 m-b-10">
                <Field
                    v-model="user.email"
                  name="email"
                  type="email"
                  placeholder="Password"
                  class="sizefull s-text7 p-l-22 p-r-22"
                  :disabled="isLoading"
                />
              </div>

              <ErrorMessage name="" class="text-danger m-b-20 d-block" />
                  <div class="col-md-6">
                    <label>Gender</label>
                  </div>
              <div class="bo4 of-hidden size15 m-b-10">
                <Field
                    v-model="user.gender"
                  name="gender"
                  type="text"
                  placeholder="Password"
                  class="sizefull s-text7 p-l-22 p-r-22"
                  :disabled="isLoading"
                />
              </div>

              <ErrorMessage name="" class="text-danger m-b-20 d-block" />
                  <div class="col-md-6">
                    <label>Address</label>
                  </div>
              <div class="bo4 of-hidden size15 m-b-10">
                <Field
                  v-model="user.address"
                  name="address"
                  type="text"
                  placeholder="Confirm Password"
                  class="sizefull s-text7 p-l-22 p-r-22"
                  :disabled="isLoading"
                />
              </div>

              <ErrorMessage
                name="confirm_password"
                class="text-danger m-b-20 d-block"
              />
                  <div class="col-md-6">
                    <label>Phone</label>
                  </div>
              <div class="bo4 of-hidden size15 m-b-10">
                <Field
                  v-model="user.phone"
                  name="phone"
                  type="text"
                  placeholder="Confirm Password"
                  class="sizefull s-text7 p-l-22 p-r-22"
                  :disabled="isLoading"
                />
              </div>

              <ErrorMessage
                name="confirm_password"
                class="text-danger m-b-20 d-block"
              />
            </Form>
          
          
          </div>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
import {mapState} from "vuex";
import {Form, Field, ErrorMessage} from "vee-validate";
import * as yup from "yup";
import * as Yup from "yup";

export default {
  name: "ProfilePage",
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
      userr: {
        full_name: '',
        email: '',
        password: ''
      }
    };
  },
  computed: {
    ...mapState("users", ["user"])
  },
  methods: {
  }
};
</script>

<style scoped>
body{
  background: -webkit-linear-gradient(left, #3931af, #00c6ff);
}
.bo4 {
  border : none;
}
.emp-profile{
  margin-bottom: 3%;
  border-radius: 0.5rem;
  background: #fff;
}
.profile-img{
  text-align: center;
}
.profile-img img{
  width: 350px;
  height: 300px;
}
.profile-img .file {
  position: relative;
  overflow: hidden;
  margin-top: -20%;
  width: 70%;
  border: none;
  border-radius: 0;
  font-size: 15px;
  background: #212529b8;
}
.profile-img .file input {
  position: absolute;
  opacity: 0;
  right: 0;
  top: 0;
}
.profile-head h3{
  margin-top: 40px;
  color: #333;
}
.profile-head h6{
  color: #0062cc;
}
.profile-edit-btn{
  border: none;
  border-radius: 1.5rem;
  width: 70%;
  padding: 2%;
  font-weight: 600;
  color: #6c757d;
  cursor: pointer;
}
.proile-rating{
  font-size: 12px;
  color: #818182;
  margin-top: 5%;
}
.proile-rating span{
  color: #495057;
  font-size: 15px;
  font-weight: 600;
}
.profile-head .nav-tabs{
  margin-bottom:5%;
}
.profile-head .nav-tabs .nav-link{
  font-weight:600;
  border: none;
}
.profile-head .nav-tabs .nav-link.active{
  border: none;
  border-bottom:2px solid #0062cc;
}
.profile-work{
  padding: 14%;
  margin-top: -15%;
}
.profile-work p{
  font-size: 12px;
  color: #818182;
  font-weight: 600;
  margin-top: 10%;
}
.profile-work a{
  text-decoration: none;
  color: #495057;
  font-weight: 600;
  font-size: 14px;
}
.profile-work ul{
  list-style: none;
}
.profile-tab label{
  font-weight: 600;
}
.profile-tab p{
  font-weight: 600;
  color: #0062cc;
}
</style>