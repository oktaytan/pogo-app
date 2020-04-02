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
    <a-col
      :xs="{ span: 24 }"
      :md="{ span: 12 }"
      class="content"
      :style="{
        background: GET_THEME
          ? GET_COLORS.dark.background
          : GET_COLORS.light.background,
        color: GET_THEME
          ? GET_COLORS.dark.textPrimary
          : GET_COLORS.light.textPrimary
      }"
    >
      <!-- Yeni kullanıcı kayıt formu -->
      <a-form
        id="customForm"
        :form="form"
        class="custom_form"
        style="margin-top: 4rem;"
        @submit="handleSubmit"
      >
        <div class="form_header animated fadeInDown">
          <img src="../assets/logo.svg" alt="" />
          <span
            :style="{
              color: GET_THEME
                ? GET_COLORS.dark.textPrimary
                : GET_COLORS.light.textPrimary
            }"
            >Pogo' ya hoşgeldin</span
          >
        </div>
        <div class="animated fadeInRight">
          <!-- Kullanıcı adı girme alanı -->
          <a-form-item>
            <a-input
              size="large"
              v-decorator="[
                'username',
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
              :style="{
                background: GET_THEME
                  ? GET_COLORS.dark.backgroundLight
                  : GET_COLORS.light.backgroundLight,
                color: GET_THEME
                  ? GET_COLORS.dark.textPrimary
                  : GET_COLORS.light.textPrimary
              }"
            />
          </a-form-item>
          <!-- Kullanıcı email adresi girme alanı -->
          <a-form-item>
            <a-input
              size="large"
              v-decorator="[
                'email',
                {
                  rules: [
                    {
                      required: true,
                      message: 'Lütfen bir email giriniz!'
                    },
                    {
                      type: 'email',
                      message: 'Lütfen geçerli bir email giriniz!'
                    }
                  ]
                }
              ]"
              placeholder="Email"
              :style="{
                background: GET_THEME
                  ? GET_COLORS.dark.backgroundLight
                  : GET_COLORS.light.backgroundLight,
                color: GET_THEME
                  ? GET_COLORS.dark.textPrimary
                  : GET_COLORS.light.textPrimary
              }"
            />
          </a-form-item>
          <!-- Kullanıcı şifre girme alanı -->
          <a-form-item>
            <a-input
              size="large"
              v-decorator="[
                'password',
                {
                  rules: [
                    { required: true, message: 'Litfen şifrenizi giriniz!' },
                    {
                      validator: validateToNextPassword
                    }
                  ]
                }
              ]"
              type="password"
              placeholder="Şifre"
              :style="{
                background: GET_THEME
                  ? GET_COLORS.dark.backgroundLight
                  : GET_COLORS.light.backgroundLight,
                color: GET_THEME
                  ? GET_COLORS.dark.textPrimary
                  : GET_COLORS.light.textPrimary
              }"
            />
          </a-form-item>
          <!-- Kullanıcı şifre tekrar girme alanı -->
          <a-form-item>
            <a-input
              size="large"
              v-decorator="[
                'password2',
                {
                  rules: [
                    {
                      required: true,
                      message: 'Litfen şifrenizi tekrar giriniz!'
                    },
                    {
                      validator: compareToFirstPassword
                    }
                  ]
                }
              ]"
              type="password"
              placeholder="Şifre Tekrar"
              :style="{
                background: GET_THEME
                  ? GET_COLORS.dark.backgroundLight
                  : GET_COLORS.light.backgroundLight,
                color: GET_THEME
                  ? GET_COLORS.dark.textPrimary
                  : GET_COLORS.light.textPrimary
              }"
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
            <span
              :style="{
                color: GET_THEME
                  ? GET_COLORS.dark.textPrimary
                  : GET_COLORS.light.textPrimary
              }"
            >
              hesabın varsa
            </span>
            <router-link to="/login">
              giriş yap!
            </router-link>
          </a-form-item>
        </div>
      </a-form>
      <div class="content_footer">
        {{ `&copy; ${new Date().getFullYear()} POGO, Inc.` }}
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
    // Register formu oluşturuluyor ( ant desing vue formu )
    this.form = this.$form.createForm(this, { name: "registerForm" });
  },
  computed: {
    ...mapGetters(["GET_THEME", "GET_COLORS"])
  },
  methods: {
    ...mapActions(["USER_REGISTER"]),
    // Şifreler birbiri ile eşleşiyor mu diye kontrol sağlanıyor
    compareToFirstPassword(rule, value, callback) {
      const form = this.form;
      if (value && value !== form.getFieldValue("password")) {
        callback("Şifreler birbiri ile eşleşmiyor!");
      } else {
        callback();
      }
    },
    // Şifre karakter sayısı kontrolü yapılıyor
    validateToNextPassword(rule, value, callback) {
      const form = this.form;
      if (value && this.confirmDirty) {
        form.validateFields(["confirm"], { force: true });
      } else if (value && value.length < 4) {
        callback("Şifre en az 4 karakter olmalıdır!");
      } else {
        callback();
      }
    },
    handleSubmit(e) {
      e.preventDefault();

      this.form.validateFields((err, values) => {
        if (!err) {
          // Form bilgilerinde hata yoksa kullanıcı kaydı için action tetikleniyor
          this.USER_REGISTER(values)
            .then(res => {
              if (!res.error && res.user_added) {
                // Register işlemi başarılı ise kullanıcı login sayfasına yönderiliyor
                this.$router.push("/login");
              } else {
                this.$message.error("Bir hata oluştu!");
                this.$message.info("Sayfayı yenileyip tekrar deneyiniz.");
              }
            })
            .catch(err => {
              if (err) {
                this.$message.error("Bir hata oluştu!");
                this.$message.info("Sayfayı yenileyip tekrar deneyiniz.");
              }
            });
        }
      });
    }
  }
};
</script>
