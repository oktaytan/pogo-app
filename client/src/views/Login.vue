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
      <!-- Login formu -->
      <a-form
        id="customForm"
        :form="form"
        class="custom_form"
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
            >Paylaşıma katıl</span
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
          <!-- Kullanıcı şifre girme alanı -->
          <a-form-item>
            <a-input
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
          <!-- Beni hatırla seçim alanı -->
          <a-form-item>
            <a-checkbox
              v-decorator="[
                'remember',
                {
                  valuePropName: 'checked',
                  initialValue: false
                }
              ]"
              :style="{
                color: GET_THEME
                  ? GET_COLORS.dark.textPrimary
                  : GET_COLORS.light.textPrimary
              }"
            >
              Beni Hatırla
            </a-checkbox>
            <a-button
              type="primary"
              html-type="submit"
              class="form_button"
              :loading="loading"
              size="large"
            >
              Giriş yap
            </a-button>
            <span
              :style="{
                color: GET_THEME
                  ? GET_COLORS.dark.textPrimary
                  : GET_COLORS.light.textPrimary
              }"
              >ya da Pogo' ya</span
            >
            <router-link to="/register">
              kayıt ol!
            </router-link>
          </a-form-item>
        </div>
      </a-form>
      <div class="content_footer">
        {{ `&copy; ${new Date().getFullYear()} POGO, Inc.` }}
      </div>
    </a-col>
    <Loader v-if="loading" />
  </a-row>
</template>

<script>
import Loader from "../components/Loader";
import { mapGetters, mapActions } from "vuex";

export default {
  name: "Login",
  components: {
    Loader
  },
  data: () => ({
    loading: false
  }),
  beforeCreate() {
    // Login formu oluşturuluyor ( ant desing vue formu )
    this.form = this.$form.createForm(this, {
      name: "loginForm",
      mapPropsToFields: () => {
        return {
          // Kullanıcı adı localStorage' dan alınıp ilgili alana yazılıyor
          username: this.$form.createFormField({
            value: localStorage.getItem("pogo_login")
              ? JSON.parse(localStorage.getItem("pogo_login")).username
              : ""
          }),
          // Beni Hatırla seçemi localStorage' dan alınıp ilgili alana yazılıyor
          remember: this.$form.createFormField({
            value: localStorage.getItem("pogo_login")
              ? JSON.parse(localStorage.getItem("pogo_login")).remember
              : false
          })
        };
      }
    });
  },
  computed: {
    ...mapGetters(["GET_THEME", "GET_COLORS"])
  },
  methods: {
    ...mapActions(["USER_LOGIN"]),
    // login submit
    handleSubmit(e) {
      e.preventDefault();
      this.form.validateFields((err, values) => {
        // Hata yoksa login isteği yapılıyor
        if (!err) {
          // Eğer beni hatırla işaretliyse
          // localStorage' e kullanıcı adı kaydı yapılıyor
          if (values.remember) {
            localStorage.setItem(
              "pogo_login",
              JSON.stringify({
                username: values.username,
                remember: values.remember
              })
            );
          } else {
            // Eğer beni hatırla işaretli değilse
            // localStorage' den kullanıcı adı kaydı siliniyor
            localStorage.setItem(
              "pogo_login",
              JSON.stringify({
                username: "",
                remember: values.remember
              })
            );
          }

          const user = {
            username: values.username,
            password: values.password
          };

          // Login yapacak action tetikleniyor
          this.USER_LOGIN(user)
            .then(res => {
              if (res.token) {
                this.loading = true;
                this.$message.loading("Giriş yapılıyor", 0.5).then(() => {
                  this.loading = false;
                  this.$router.push("/");
                });
              }
            })
            .catch(err => {
              this.$message.error(err.message);
            });
        }
      });
    }
  }
};
</script>


