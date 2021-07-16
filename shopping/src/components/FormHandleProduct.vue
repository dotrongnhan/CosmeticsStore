<template>
  <div id="myModal" class="modal">
    <div class="modal-content">
      <h1>PRODUCT</h1>
      <div style="justify-content: center;display: flex" class="container p-t-66 p-b-60">
        <Form @submit.prevent="submit"  style="min-width: 65%;" :validation-schema="schema">
          <div  class="bo4 of-hidden size15 m-b-5">
            <Field
                name="name"
                type="text"
                placeholder="Product's name"
                class="sizefull s-text7 p-l-22 p-r-22"
            />
          </div>
          <ErrorMessage name="name" class="fs-18 text-danger d-block" />

          <div class="bo4 of-hidden size15 m-b-5">
            <Field
                name="image"
                type="text"
                placeholder="Product's image"
                class="sizefull s-text7 p-l-22 p-r-22"
            />
          </div>
          <ErrorMessage name="image" class="fs-18 text-danger d-block" />
          <div class="row">
            <div class="col-6">
              <div class="bo4 of-hidden size15 m-b-5">
                <Field
                    name="price"
                    type="number"
                    placeholder="Product's price"
                    class="sizefull s-text7 p-l-22 p-r-22"
                />
              </div>
              <ErrorMessage name="price" class="fs-18 text-danger d-block" />

            </div>
            <div class="col-6">
              <div class="bo4 of-hidden size15 m-b-5">
                <Field
                    name="numberAvailable"
                    type="number"
                    placeholder="Product's number available"
                    class="sizefull s-text7 p-l-22 p-r-22"
                />
              </div>
              <ErrorMessage name="numberAvailable" class="fs-18 text-danger d-block" />

            </div>
          </div>



        </Form>

      </div>
      <div class="button">
        <button @click="yesConfirm">YES</button>
        <button @click="noConfirm">NO</button>
      </div>
    </div>

  </div>
</template>

<script>
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";

export default {
  name: "FormHandleProduct",
  components:{
    Form,
    Field,
    ErrorMessage
  },
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
    }
  }
}
</script>

<style scoped>
.button {
  display: flex;
  justify-content: space-around;
}
[type=button]:not(:disabled), [type=reset]:not(:disabled), [type=submit]:not(:disabled), button:not(:disabled) {
  border-radius: 10px;
  cursor: pointer;
  height: 40px;
  width: 400px;
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
  width: 80%;
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