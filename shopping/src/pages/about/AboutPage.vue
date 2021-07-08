<template>
  <div>
    <!-- Title Page -->
    <section
      class="bg-title-page p-t-40 p-b-50 flex-col-c-m"
      :style="{
        backgroundImage:
          'url(' + require('@/assets/bd6ba-cosmetics.jpeg') + ')',
      }"
    >
    </section>

    <!-- content page -->
    <section class="bgwhite p-t-66 p-b-38">
      <div class="container">
        <div class="row">
          <div class="col-md-4 p-b-30">
            <div class="hov-img-zoom">
              <img src="https://www.brandinginasia.com/wp-content/uploads/2020/02/Kasumi-Ishikawa-Beauty-is-No-Competition-Branding-in-Asia.jpg" alt="IMG-ABOUT" />
            </div>
          </div>

          <div class="col-md-8 p-b-30">
            <h3 class="m-text26 p-t-15 p-b-16">Our story</h3>

            <p class="p-b-28">
              History
              SK-II was developed in the 1970s by Japanese scientists investigating the use of more naturally derived ingredients after it was observed that elderly workers at a sake brewery had smooth, clear and young-looking hands from years of submerging them in the fermented yeast. A yeast extract that the company calls pitera was eventually isolated for cosmetic use and the brand was launched in the early 1980s. P&G acquired the brand with its purchase of Max Factor in 1991 and expanded its sale from Japan to Taiwan, Korea, Hong Kong, in 2000 to the UK, and subsequently to the US, restricting initial sales to a few high-end stores where "consultants" introduced customers to the brand. As of 2018, it is also sold in Australia, Indonesia, Malaysia, Singapore, Spain, and Thailand.
            </p>

            <div class="bo13 p-l-29 m-l-9 p-b-10">
              <p class="p-b-11">
                Controversies in China
                In April 2005, P&G agreed to pay a 200,000 yuan fine after a customer in China alleged inaccurate advertising of an SK-II product and that it had caused an allergic reaction. In September 2006, traces of neodymium and chromium were found in SK-II products sold in China. Both are banned in cosmetics, and can cause allergic dermatitis and eczema. P&G China withdrew several products, while requiring consumers to sign a "safe product" agreement releasing P&G from liability in order to return those they had bought. By the end of that week, P&G had suspended sales of the brand altogether in China and closed its sales counters there, instituting a hotline for refund applications. SK-II sales were also suspended for approximately two weeks in South Korea, resuming after authorities announced the products were safe. P&G announced it would resume selling the products in China in late 2006 after regulators declared that the trace elements of these substances in cosmetics were not likely to be injurious to human health.
              </p>

              <span class="s-text7"> - Steve Job’s </span>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: "AboutPage",
};
</script>