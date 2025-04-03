<template>
  <nav
    class="navbar has-shadow is-fixed-top"
    role="navigation"
    aria-label="main navigation"
  >
    <div class="container" >
      <div class="navbar-brand">
        
        <a
          :class="{ 'is-active': displayStore.isEnable }"
          class="navbar-burger burger"
          data-target="navbarBasic"
          @click="toggleNav"
        >
          <span aria-hidden="true" />
          <span aria-hidden="true" />
          <span aria-hidden="true" />
        </a>
      </div>

       
    </div>
  </nav>
</template>

<script setup>
 const userStore = useUserStore();
 const configStore = useConfigStore();
 const displayStore = useDisplayStore();
 const { user } = storeToRefs(userStore);
 const { config } = storeToRefs(configStore);
console.log("displayStore", displayStore.config.navbarActive);
// const navbarActive = displayStore.config.navbarActive;

/**
 * 切换导航栏的显示状态
 *
 * @function toggleNav
 * @returns {void} 无返回值
 */
function toggleNav() {
  displayStore.increment();
  // displayStore.config.navbarActive = !displayStore.config.navbarActive;
  
}

// async function signout() {
//   if (confirm("确定退出登录吗？")) {
//     await userStore.signout();
//     useLinkTo("/");
//   }
// }
</script>

<style lang="scss" scoped>
.navbar {
  /*opacity: 0.99;*/
  /*border-bottom: 1px solid #e7edf3;*/
  background-color: var(--bg-color);

  .navbar-item {
    font-weight: 700;
  }

  .publish {
    color: var(--text-color);
    background-color: #3174dc;
    width: 100px;
    &:hover {
      color: var(--text-color);
      background-color: #4d91fa;
    }
  }

  .login-btn {
    height: 32px;
    border-color: #000; // TODO
    &:hover {
      color: var(--text-color3);
      border-color: var(--text-color3);
    }
  }
}

  .is-show {
    display: none;
  }

.user-menus {
  .navbar-link {
    display: flex;
    align-items: center;

    .user-menus-nickname {
      margin-left: 5px;
      padding: 0 4px;
      font-size: 14px;
      color: var(--text-color);

      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
    }
  }
  .navbar-dropdown {
    border: 1px solid var(--border-color);

    a {
      display: flex;
      align-items: center;
      // padding: 8px 16px;
      img {
        width: 20px;
        height: 20px;
      }
      span {
        margin-left: 10px;
        width: 56px;
        height: 20px;
        font-size: 14px;
        font-weight: 400;
        line-height: 20px;
      }
    }
  }
}
</style>
