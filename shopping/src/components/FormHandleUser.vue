<template>
  <div id="myModal" class="modal">
    <div class="modal-content">
        <div class="row">      
          <div class="profile-img">
            <!-- <img :src="user.avatar" alt="" v-show="user.avatar!=''"/>
            <img src="https://i1.wp.com/www.baytekent.com/wp-content/uploads/2016/12/facebook-default-no-profile-pic1.jpg?ssl=1" alt="" v-show="user.gender==='Male'&&user.avatar==''"/>
            <img src="https://i.pinimg.com/originals/e0/5e/06/e05e061bfc4c6e7799de4aac293a8223.png" alt="" v-show="user.gender==='Female'&&user.avatar==''"/> -->
            <img src="https://i1.wp.com/www.baytekent.com/wp-content/uploads/2016/12/facebook-default-no-profile-pic1.jpg?ssl=1" />
          </div>
          <h1>USER NAME</h1>
        </div>  
      <div style="display: flex;justify-content: center;" class="container p-t-66 p-b-60">
       <form>

      <div class="primary-info">    
       <div>  
        <div class="mb-3">
          <label class="form-label fs-4">Email</label>
          <p>Now</p>
        </div>
        <div class="mb-3">
          <label class="form-label fs-4">Phone</label>
          <p>Now</p>
        </div>
       </div>
       <div> 
        <div class="mb-3">
          <label class="form-label fs-4">Gender</label>
          <p>Now</p>
        </div>
        <div class="mb-3">
          <label class="form-label fs-4">Date of Birth</label>
          <p>Now</p>
        </div>
       </div>
      </div>  
        <div class="mb-3" style="float:left;">
          <label class="form-label fs-4">Address</label>
          <p>Now</p>
        </div>                
      </form>   
      </div>
      <div class="button">
        <button class="close-btn" @click="closeForm">Close</button>
      </div>
    </div>

  </div>
</template>

<script>
import { mapState} from "vuex";
// import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";

export default {
  name: "FormHandleUser",
  // components:{
  //   Form,
  //   Field,
  //   ErrorMessage
  // },
  data() {
    const schema = yup.object().shape({
      name: yup
          .string()
          .required("Product's name is required"),
      image: yup
          .string()
          .required("Product's image is required"),
      price: yup
          .number()
          .required("Product's price is required"),
      hot: yup
          .number(),
      category: yup
          .number()
          .required("Product's category is required"),
      branch: yup
          .number()
          .required("Product's category is required"),
      numberAvailable: yup
          .number()
          .required("Product's number available is required")
    })
    return {
      schema,
    }
  },
  methods: {
    noConfirm: function () {
      this.$emit("confirmDeleteItem", false)
    },
    yesConfirm: function () {
      this.$emit("confirmDeleteItem", true)
    },
    closeForm() {
      this.$store.commit("users/displayForm");
    }
  },
  computed: {
    ...mapState("users", ["isDisplayForm"]),
  }
}
</script>

<style scoped>
.button {
  display: flex;
  justify-content: space-around;
}
.close-btn:hover {
    color: red;
}
[type=button]:not(:disabled), [type=reset]:not(:disabled), [type=submit]:not(:disabled), button:not(:disabled) {
  border-radius: 10px;
  cursor: pointer;
  height: 40px;
  width: 400px;
}
.profile-img{
  text-align: center;
}
.profile-img img{
  width: 350px;
  height: 300px;
}
.primary-info {
  display: grid;
  grid-template-columns: auto auto;
  grid-gap: 200px;
}

.modal {
  text-align: center;
  font-size: 25px;
  display: block;
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgb(0,0,0);
  background-color: rgba(0,0,0,0.4);
}

.modal-content {
  background-color: #fefefe;
  margin: 15% auto;
  padding: 20px;
  border: 1px solid #888;
  width: 50%;
}

/* The Close Button */
.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: black;
  text-decoration: none;
  cursor: pointer;
}
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
</style>