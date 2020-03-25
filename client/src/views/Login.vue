<template>
  <a-row :gutter="0" style="height: 100vh; overflow: hidden;">
    <a-col
      :xs="{ span: 0 }"
      :md="{ span: 12 }"
      style="height: 100%; position: relative; overflow: hidden"
    >
      <img
        style="height: 100%; position: absolute; top: 0; left: 0; right: 0"
        src="../assets/login-bg.png"
        alt=""
      />
      <div class="bg_desc">
        <div class="animated fadeInLeft">
          <p><a-icon type="search" />Paylaşımları gör</p>
          <p><a-icon type="check" />Yorum yap</p>
          <p><a-icon type="highlight" />Sende katıl</p>
        </div>
      </div>
    </a-col>
    <a-col :xs="{ span: 24 }" :md="{ span: 12 }" class="content">
      <a-form
        id="customForm"
        :form="form"
        class="custom_form"
        @submit="handleSubmit"
      >
        <div class="form_header">
          <img src="../assets/logo.svg" alt="" />
          <span>Paylaşıma katıl</span>
        </div>
        <div class="animated fadeInRight">
          <a-form-item>
            <a-input
              size="large"
              v-decorator="[
                'userName',
                {
                  rules: [
                    {
                      required: true,
                      message: 'Lütfen kullanıcı adınızı giriniz!'
                    }
                  ]
                }
              ]"
              placeholder="Kullanıcı Adı"
            />
          </a-form-item>
          <a-form-item>
            <a-input-password
              size="large"
              v-decorator="[
                'password',
                {
                  rules: [
                    { required: true, message: 'Litfen şifrenizi giriniz!' }
                  ]
                }
              ]"
              type="password"
              placeholder="Şifre"
            />
          </a-form-item>
          <a-form-item>
            <a-checkbox
              v-decorator="[
                'remember',
                {
                  valuePropName: 'checked',
                  initialValue: false
                }
              ]"
            >
              Beni Hatırla
            </a-checkbox>
            <a-button
              type="primary"
              html-type="submit"
              class="form_button"
              size="large"
            >
              Giriş yap
            </a-button>
            ya da Pogo' ya
            <router-link to="/register">
              kayıt ol!
            </router-link>
          </a-form-item>
        </div>
      </a-form>
      <div class="content_footer">
        {{ `&copy; ${new Date().getFullYear()} POGO, Inc.`  }}
      </div>
    </a-col>
  </a-row>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "Login",
  data: () => ({}),
  beforeCreate() {
    this.form = this.$form.createForm(this, { name: "loginForm" });
  },
  mounted() {},
  computed: {
    ...mapGetters(["getColors"])
  },
  methods: {
    handleSubmit(e) {
      e.preventDefault();
      this.form.validateFields((err, values) => {
        if (!err) {
          console.log("Received values of form: ", values);
        }
      });
    }
  }
};
</script>

<style scoped></style>
