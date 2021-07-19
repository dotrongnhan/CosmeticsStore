<template>
           <div id="snackbar">Update successfully</div> 
  <div class="container emp-profile">
    <form method="post">
      <div class="row">
        <div class="col-md-4 pb-2">
          <div class="profile-img">
            <img :src="user.avatar" alt="" v-show="user.avatar!=''"/>
            <img src="https://i1.wp.com/www.baytekent.com/wp-content/uploads/2016/12/facebook-default-no-profile-pic1.jpg?ssl=1" alt="" v-show="user.gender==='Male'&&user.avatar==''"/>
            <img src="https://i.pinimg.com/originals/e0/5e/06/e05e061bfc4c6e7799de4aac293a8223.png" alt="" v-show="user.gender==='Female'&&user.avatar==''"/>
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
                <button class="nav-link active" id="nav-profile-tab" data-bs-toggle="tab" data-bs-target="#nav-profile" type="button" role="tab" aria-controls="nav-profile" aria-selected="true" @click="toAbout">About</button>
              </li>
              <li class="nav-item">
                <!-- <a class="nav-link " id="home-tab" data-toggle="tab" href="#home" role="tab" aria-controls="home" aria-selected="true">Security</a> -->
                <button class="nav-link" id="nav-security-tab" data-bs-toggle="tab" data-bs-target="#nav-security" type="button" role="tab" aria-controls="nav-security" aria-selected="false" @click="toSecurity">Security</button>
              </li>
            </ul>
          </div>
        </div>
        <div class="col-md-2">
        </div>
      </div>

      <!-- Profile  -->
      <div class="row" id="profile" v-show="status==='about'">
        <div class="col-md-8">
          <div class="tab-content profile-tab">
            <Form @submit="onSubmit" :validation-schema="schemaProfile" style="margin-left: 390px; margin-top: -100px; width : 74%;">
              <div>
                <div
                    style = "width: 30%; margin-left: 390px"
                    @click="editProfile"
                    class="flex-c-m size2 bg1 bo-rad-23 hov1 m-text3 trans-0-4 m-t-20 edit-button"
                  :class="{ disabled: isLoading }"
                  :disabled="isLoading"
                  v-show="!onEdit"
                >
                  <span v-show="isLoading" data-loader="ball-scale"></span>
                  Edit Profile
                </div>
                <div
                    style = "width: 30%; margin-left: 390px"
                    @click="removeEditProfile"
                    class="flex-c-m size2 bg1 bo-rad-23 hov1 m-text3 trans-0-4 m-t-20 edit-button"
                  :class="{ disabled: isLoading }"
                  :disabled="isLoading"
                  v-show="onEdit"
                >
                  <span v-show="isLoading" data-loader="ball-scale"></span>
                  Cancel
                </div>
              </div>
                <div v-show="onEdit">              
                  <div class="col-md-6">
                    <label>Upload Avatar Here: </label>
                  </div>
                  <div class="bo4 of-hidden size15 m-b-10">
                    <Field
                      v-model="user.avatar"
                      name="avatar"
                      type="text"
                      placeholder='Paste your photo link here ...'
                      class="sizefull s-text7 p-l-22 p-r-22"
                      :disabled="isLoading"
                    />
                  </div>
              </div>

                  <div class="col-md-6">
                    <label>Full name</label>
                  </div>
              <div class="bo4 of-hidden size15 m-b-10">
                <Field
                    v-model="user.full_name"
                  name="full_name"
                  type="text"
                  placeholder='Your name ...'
                  class="sizefull s-text7 p-l-22 p-r-22"
                  :disabled="isLoading"
                />
              </div>

              <ErrorMessage
                name="full_name"
                class="text-danger m-b-20 d-block"
              />

            <div class="dob-gender" style="display: flex; justify-content: space-between;">
              <div style="width: 100%;">
                  <div class="col-md-6">
                    <label>Date of Birth</label>
                  </div>
              <div class="bo4 of-hidden size15 m-b-10" >
                <Field
                  v-model="user.date_of_birth"
                  name="date_of_birth"
                  type="date"
                  placeholder="When were you born?..."
                  class="sizefull s-text7 p-l-22 p-r-22"
                  :disabled="isLoading"
                  style="width: 70%;"
                />
              </div>
              </div>

              <div style="width: 100%;">
                  <div class="col-md-6">
                    <label>Gender</label>
                  </div>
              <div class="bo4 of-hidden size15 m-b-10" >
                <Field
                  v-model="user.gender"
                  as="select"
                  name="gender"
                  id="genderSelect"
                  type="text"
                  placeholder="Password"
                  class="sizefull s-text7 p-l-22 p-r-22"
                  :disabled="isLoading"
                  style="width: 50%;"
                >
                    <option value="Male">Male</option>
                    <option value="Female">Female</option>
                </Field>
              </div>
              </div>
            </div>  

                  <div class="col-md-6">
                    <label>Address</label>
                  </div>
              <div class="bo4 of-hidden size15 m-b-10">
                <Field
                  v-model="user.address"
                  name="address"
                  type="text"
                  placeholder="Where do you live?... "
                  class="sizefull s-text7 p-l-22 p-r-22"
                  :disabled="isLoading"
                />
              </div>

                  <div class="col-md-6">
                    <label>Phone</label>
                  </div>
              <div class="bo4 of-hidden size15 m-b-10">
                <Field
                  v-model="user.phone"
                  name="phone"
                  type="text"
                  placeholder="Your phone number ..."
                  class="sizefull s-text7 p-l-22 p-r-22"
                  :disabled="isLoading"
                />
              </div>

              <ErrorMessage
                name="phone"
                class="text-danger m-b-20 d-block"
              />
              <button
                    style="margin-left: 195px; width: 30%;"
                    @click="removeEditProfile(); myFunction();"
                    class="flex-c-m size2 bg1 bo-rad-23 hov1 m-text3 trans-0-4 m-t-20"
                  :class="{ disabled: isLoading }"
                  :disabled="isLoading"
                  v-show="onEdit"
                >
                  <span v-show="isLoading" data-loader="ball-scale"></span>
                  Confirm
                </button>
            </Form>
          
          
          </div>
        </div>
      </div>

      <!-- Password  -->
      <div class="row" id="password" v-show="status==='security'">
        <div class="col-md-8">
          <div class="tab-content profile-tab" id="myTabContent">
          
            <Form @submit="onSubmit1" :validation-schema="schemaPassword" style="margin-left: 390px; margin-top: -100px; width : 74%;">
                          <p class="m-b-20 text-center text-danger" v-show="!isChangePassSuccess">
                            {{ changePassMessage }}
                          </p>
                          <p class="m-b-20 text-center text-primary" v-show="isChangePassSuccess">
                            {{ changePassMessage }}
                          </p>
              <div class="col-md-6">
                    <label>Current Password</label>
              </div>
              <div class="bo4 of-hidden size15 m-b-10">
                <Field
                  v-model="userr.current_password"
                  name="current_password"
                  type="password"
                  placeholder='Your current password...'
                  class="sizefull s-text7 p-l-22 p-r-22 password-field"
                  :disabled="isLoading"
                />
              </div>
              <ErrorMessage
                name="current_password"
                class="text-danger m-b-20 d-block"
              />

              <div class="col-md-6">
                    <label>New Password</label>
                  </div>
              <div class="bo4 of-hidden size15 m-b-10">
                <Field
                  v-model="userr.new_password"
                  name="new_password"
                  type="password"
                  placeholder="Your new password..."
                  class="sizefull s-text7 p-l-22 p-r-22 password-field"
                  :disabled="isLoading"
                />
              </div>
              <ErrorMessage
                name="new_password"
                class="text-danger m-b-20 d-block"
              />

                 <div class="col-md-6">
                    <label>Confirm New Password</label>
                  </div>
              <div class="bo4 of-hidden size15 m-b-10">
                <Field
                  name="confirm_password"
                  type="password"
                  placeholder="Confirm your new password..."
                  class="sizefull s-text7 p-l-22 p-r-22 password-field"             
                  :disabled="isLoading"
                />
              </div>
              <ErrorMessage
                name="confirm_password"
                class="text-danger m-b-20 d-block"
              />

              <button
                    style="margin-left: 150px; width: 40%;"
                    class="flex-c-m size2 bg1 bo-rad-23 hov1 m-text3 trans-0-4 m-t-20"
                  :class="{ disabled: isLoading }"
                  :disabled="isLoading"
                >
                <span v-show="isLoading" data-loader="ball-scale"></span>
                  Change Password
                </button>
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
import * as yup1 from "yup";
import * as Yup from "yup";
export default {
  name: "ProfilePage",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  data() {
    const schemaPassword = yup.object().shape({
      current_password: yup
        .string()
        .required("Password is required!")
        ,
      new_password: yup
        .string()
        .required("Password is required!")
        .min(6, "Password must be at least 6 characters!")
        .max(40, "Password must be maximum 40 characters!"),
      confirm_password: yup
      .string()
      .oneOf([Yup.ref('new_password'), null], 'Passwords must match')
    });

    const schemaProfile = yup1.object().shape({
      full_name: yup1
        .string()
        .required("Your name is required!")
        .min(1),
      phone: yup1
        .number()
        .typeError("That doesn't look like a phone number")
        .positive("A phone number can't start with a minus")
        .integer("A phone number can't include a decimal point")
        .min(8)
    });

    return {
      onEdit: false,
      isLoading: false,
      message: "",
      status : "about",
      schemaPassword,
      schemaProfile,
      userr: {
          current_pasword: '',
          new_password: ''
      }
    };
  },
  computed: {
    ...mapState("users", ["user", "isUpdateSuccess", "updateMessage", "isChangePassSuccess", "changePassMessage"])
  },
  methods: {
      toSecurity() {
          document.getElementById("nav-security-tab").className = "nav-link active";
          document.getElementById("nav-profile-tab").className = "nav-link";
          this.status = "security";
         let elements = document.getElementsByClassName("bo4");
            for(let i = 0; i < elements.length; i++) {
                  elements[i].style.border = '0.5px solid';
            }
      },
      toAbout() {
          document.getElementById("nav-security-tab").className = "nav-link";
          document.getElementById("nav-profile-tab").className = "nav-link active";
          this.status = "about";
          let elements = document.getElementsByClassName("bo4");
            for(let i = 0; i < elements.length; i++) {
                  elements[i].style.border = 'none';
            }
      },  
      openPage(pageName,elmnt,color) {
        var i, tabcontent, tablinks;
        tabcontent = document.getElementsByClassName("tabcontent");
        for (i = 0; i < tabcontent.length; i++) {
          tabcontent[i].style.display = "none";
        }
        tablinks = document.getElementsByClassName("tablink");
        for (i = 0; i < tablinks.length; i++) {
          tablinks[i].style.backgroundColor = "";
        }
        document.getElementById(pageName).style.display = "block";
        elmnt.style.backgroundColor = color;
      },
    myFunction() {
          var x = document.getElementById("snackbar");
        if(this.isSecurity) {       
          x.className = "show";
          setTimeout(function(){ x.className = x.className.replace("show", ""); }, 3000);
        }
    },
    editProfile() {
     let elements = document.getElementsByClassName("bo4");
      for(let i = 0; i < elements.length; i++) {
          elements[i].style.border = '0.5px groove';
      }
      this.onEdit = true;
    },
    removeEditProfile() {
     let elements = document.getElementsByClassName("bo4");
      for(let i = 0; i < elements.length; i++) {
          elements[i].style.border = 'none';
      }
      this.onEdit = false;
    },
    onSubmit() {
      this.$store.dispatch("users/updateProfile", {
        avatar: this.user.avatar, 
        full_name: this.user.full_name,
        date_of_birth: this.user.date_of_birth,
        gender: this.user.gender,
        address: this.user.address,
        phone: this.user.phone,
      });
      this.isLoading = false;
      if (this.isUpdateSuccess) {
        this.message = "Update successfully.";
      }
    },
    onSubmit1() {
      this.$store.dispatch("users/changePassword", {
              current_password : this.userr.current_password,
              new_password : this.userr.new_password,
      })
      this.isLoading = false;
    setTimeout(() => {  
        let elements = document.getElementsByClassName('password-field');
                for(let i = 0; i < elements.length; i++) {
                      elements[i].value = '';
                }
              }, 1800);              
      if (this.isChangePassSuccess) {
        this.message = "Change successfully.";
      }
    },
  }
};
</script>

<style scoped>
body{
  background: -webkit-linear-gradient(left, #3931af, #00c6ff);
}

#genderSelect {
  border: none;

}
.edit-button:hover {
  cursor: pointer;
} 

.bo4 {
  border : none;
}
.sizefull:hover {
  cursor: default;
}
.nav-link {
  color: black;
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
#snackbar {
  visibility: hidden;
  min-width: 250px;
  margin-left: -125px;
  background-color: #333;
  color: #fff;
  text-align: center;
  border-radius: 2px;
  padding: 16px;
  position: fixed;
  z-index: 1;
  left: 50%;
  bottom: 330px;
  font-size: 17px;
}

#snackbar.show {
  visibility: visible;
  -webkit-animation: fadein 0.5s, fadeout 0.5s 2.5s;
  animation: fadein 0.5s, fadeout 0.5s 2.5s;
}

@-webkit-keyframes fadein {
  from {bottom: 0; opacity: 0;} 
  to {bottom: 330px; opacity: 1;}
}

@keyframes fadein {
  from {bottom: 0; opacity: 0;}
  to {bottom: 330px; opacity: 1;}
}

@-webkit-keyframes fadeout {
  from {bottom: 330px; opacity: 1;} 
  to {bottom: 0; opacity: 0;}
}

@keyframes fadeout {
  from {bottom: 330px; opacity: 1;}
  to {bottom: 0; opacity: 0;}
}
</style>
