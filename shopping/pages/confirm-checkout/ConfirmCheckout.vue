<template>
  <div class="container">
    <div class="py-5 text-center"></div>
    <div class="row">
      <div class="col-md-4 order-md-2 mb-4 text-secondary">
        <h4 class="d-flex justify-content-between align-items-center mb-3">
          <span class="text-muted">Your cart</span>
        </h4>
        <ul class="list-group mb-3">
          <li class="list-group-item d-flex justify-content-between lh-condensed">
            <div>
              <h6 class="my-0">Product name</h6>
              <small class="text-muted">Brief description</small>
            </div>
            <span class="text-muted">$12</span>
          </li>
          <li class="list-group-item d-flex justify-content-between lh-condensed">
            <div>
              <h6 class="my-0">Second product</h6>
              <small class="text-muted">Brief description</small>
            </div>
            <span class="text-muted">$8</span>
          </li>
          <li class="list-group-item d-flex justify-content-between lh-condensed">
            <div>
              <h6 class="my-0">Third item</h6>
              <small class="text-muted">Brief description</small>
            </div>
            <span class="text-muted">$5</span>
          </li>
          <li class="list-group-item d-flex justify-content-between bg-light">
            <div class="text-success">
              <h6 class="my-0">Promo code</h6>
              <small>EXAMPLECODE</small>
            </div>
            <span class="text-success">-$5</span>
          </li>
          <li class="list-group-item d-flex justify-content-between">
            <span>Total (USD)</span>
            <strong>$20</strong>
          </li>
        </ul>
      </div>
      <div class="col-md-8 order-md-1">
        <h4 class="mb-3">Confirm Information</h4>
            <div class="row mb-2">
              <div class="col-3 mt-1"><h6>First name : </h6></div>
              <div class="col-9"><p>TRONG</p></div>
            </div>
            <div class="row mb-2">
              <div class="col-3 mt-1"><h6>Last name : </h6></div>
              <div class="col-9"><p class="text-secondary">NHAN</p></div>
            </div>
          <div class="row mb-2">
              <div class="col-3 mt-1"><h6>Username : </h6></div>
              <div class="col-9"><p class="text-secondary">TRONG NHAN</p></div>
          </div>
          <div class="row mb-2">
              <div class="col-3 mt-1"><h6>Email : </h6></div>
              <div class="col-9"><p class="text-secondary">abc@example.com</p></div>
          </div>
          <div class="row mb-2">
              <div class="col-3 mt-1">
                <h6>Address : </h6>
              </div>
              <div class="col-9"><p class="text-secondary">abc street xyc city</p></div>
          </div>
          <h4 class="mb-4">Payment <span>Credit card</span></h4>
            <div class="row mb-2">
              <div class="col-3 mt-1">
                <h6>Name on card : </h6>
              </div>
              <div class="col-9">
                <span class="text-secondary">DO TRONG NHAN</span>
              </div>
            </div>
            <div class="row mb-2">
              <div class="col-3 mt-1">
                <h6>Credit card number : </h6>
              </div>
              <div class="col-9">
                <span class="text-secondary">3xxx xxxx xxxx xxxx</span>
              </div>
            </div>
            <div class="row mb-2">
              <div class="col-3 mt-1"><h6>Expiration : </h6></div>
              <div class="col-9"><span class="text-secondary"> 13/08</span></div>
            </div>
            <div class="mb-3 row">
              <div class="col-3"><h6>CVV</h6></div>
              <div class="col-9"><p>4523</p></div>
            </div>
          <hr class="mb-4">
          <div class="row">
            <div class="col-6">
              <button class="btn btn-primary btn-lg btn-block" type="submit">Back to edit infor</button>
            </div>
            <div class="col-6">
              <button class="btn btn-primary btn-lg btn-block" type="submit">Confirm checkout</button>
            </div>
          </div>
      </div>
    </div>

    <footer class="my-5 pt-5 text-muted text-center text-small">
      <p class="mb-1">&copy; 2017-2019 Company Name</p>
      <ul class="list-inline">
        <li class="list-inline-item"><a href="#">Privacy</a></li>
        <li class="list-inline-item"><a href="#">Terms</a></li>
        <li class="list-inline-item"><a href="#">Support</a></li>
      </ul>
    </footer>
  </div>

</template>

<script>
export default {
  name: "ConfirmCheckout"
}
</script>

<style scoped>

</style>