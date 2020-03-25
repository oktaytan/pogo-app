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
      <div class="bg_desc bg_desc--center">
        <img
          class="register_logo animated slideInUp"
          src="../assets/logo-white.svg"
          alt=""
        />
      </div>
    </a-col>
    <a-col :xs="{ span: 24 }" :md="{ span: 12 }" class="content">
      <a-form
        id="customForm"
        :form="form"
        class="custom_form"
        style="margin-top: 4rem;"
        @submit="handleSubmit"
      >
        <div class="form_header">
          <img src="../assets/logo.svg" alt="" />
          <span>Pogo' ya hoşgeldin</span>
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
            <a-input
              size="large"
              v-decorator="[
                'email',
                {
                  rules: [
                    {
                      required: true,
                      message: 'Lütfen geçerli bir email giriniz!'
                    }
                  ]
                }
              ]"
              placeholder="Email"
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
            <a-input-password
              size="large"
              v-decorator="[
                'password2',
                {
                  rules: [
                    {
                      required: true,
                      message: 'Litfen şifrenizi tekrar giriniz!'
                    }
                  ]
                }
              ]"
              type="password"
              placeholder="Şifre Tekrar"
            />
          </a-form-item>
          <a-form-item>
            <a-button
              type="primary"
              html-type="submit"
              class="form_button"
              size="large"
            >
              Kayıt Ol
            </a-button>
            hesabın varsa
            <router-link to="/login">
              giriş yap!
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
  name: "Register",
  data: () => ({}),
  beforeCreate() {
    this.form = this.$form.createForm(this, { name: "registerForm" });
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
