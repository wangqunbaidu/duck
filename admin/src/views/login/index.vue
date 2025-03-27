<template>
  <div class="login-container">
    <div class="logo">
      <img
        alt="logo"
        :src="adminLogo" width="40"
      />
      <div class="logo-text">管理系统</div>
    </div>
    <LoginBanner />
    <div class="content">
      <div class="content-inner">
        <a-spin :loading="loading" tip="Loading">
          <LoginForm />
        </a-spin>
      </div>
      <div class="footer">
        <Footer />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import logo2 from '@/assets/images/logo.png';
  import Footer from '@/components/footer/index.vue';
  import { DEFAULT_ROUTE_NAME } from '@/router/constants';
  import LoginBanner from './components/banner.vue';
  import LoginForm from './components/login-form.vue';
  

  const userStore = useUserStore();
  const router = useRouter();
  const loading = ref(true);
  const adminLogo = logo2;
  onMounted(async () => {
    try {
      const user = await userStore.info();
      if (user) {
        const { redirect, ...othersQuery } = router.currentRoute.value.query;
        router.push({
          name: (redirect as string) || DEFAULT_ROUTE_NAME,
          query: {
            ...othersQuery,
          },
        });
      }
    } finally {
      loading.value = false;
      // adminLogo.value = logo;
    }
  });
</script>

<style lang="less" scoped>
  .login-container {
    display: flex;
    height: 100vh;

    .banner {
      width: 550px;
      background: linear-gradient(163.85deg, #1d2129 0%, #00308f 100%);
      @media screen and (max-width: 1400px) {
        & {
          width: 450px !important;
        }
      }
      @media screen and (max-width: 1024px) {
        & {
          display: none !important;
        }
      }
    }

    .content {
      position: relative;
      display: flex;
      flex: 1;
      align-items: center;
      justify-content: center;
      padding-bottom: 40px;
    }

    .footer {
      position: absolute;
      right: 0;
      bottom: 0;
      width: 100%;
    }
  }

  .logo {
    position: fixed;
    top: 24px;
    left: 22px;
    z-index: 1;
    display: inline-flex;
    align-items: center;

    &-text {
      margin-right: 4px;
      margin-left: 4px;
      color: var(--color-fill-1);
      font-size: 20px;
    }
  }
</style>

<style lang="less" scoped>
  // responsive
  @media (max-width: @screen-lg) {
    .container {
      .banner {
        width: 25%;
      }
    }
  }
</style>
