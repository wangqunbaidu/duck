<template>
  <div class="dashboard">
    <a-card v-if="systemInfo" title="服务器信息">
      <div class="attr">
        <span class="attr-name">OS:</span>
        <span class="attr-value">{{ systemInfo.os }}</span>
      </div>
      <div class="attr">
        <span class="attr-name">Arch:</span>
        <span class="attr-value">{{ systemInfo.arch }}</span>
      </div>
      <div class="attr">
        <span class="attr-name">Cpu:</span>
        <span class="attr-value">{{ systemInfo.numCpu }}</span>
      </div>
      <div class="attr">
        <span class="attr-name">Go version:</span>
        <span class="attr-value">{{ systemInfo.goVersion }}</span>
      </div>
      <div class="attr">
        <span class="attr-name">Hostname:</span>
        <span class="attr-value">{{ systemInfo.hostname }}</span>
      </div>
    </a-card>

    <a-card v-if="systemInfo" title="开发团队信息">
      <div class="attr">
        <span class="attr-name">Official Site:&nbsp;</span>
        <a class="attr-value" href="/" target="_blank"
          >我们的用户端</a
        >
      </div>
      <div class="attr">
        <span class="attr-name">Dev Team:&nbsp;</span>
        <a class="attr-value" href="" target="_blank"
          >不想当牛马</a
        >
      </div>
      <div class="attr">
        <span class="attr-name">Github:&nbsp;</span>
        <a
          class="attr-value"
          href=""
          target="_blank"
          >代码地址</a
        >
      </div>
      <div class="attr">
        <span class="attr-name">Gitee:&nbsp;</span>
        <a
          class="attr-value"
          href=""
          target="_blank"
          >云管理地址</a
        >
      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
  const systemInfo = ref();
  const load = async () => {
    systemInfo.value = await axios.get('/api/admin/common/system_info');
  };
  load();
</script>

<style lang="less" scoped>
  .dashboard {
    margin: 20px;
    display: flex;
    flex-direction: column;
    row-gap: 20px;

    font-size: 14px;
    line-height: 1.8;

    .attr {
      display: flex;
      align-items: center;
      column-gap: 10px;
      font-size: 14px;
      .attr-name {
        color: var(--color-text-1);
      }
      span.attr-value {
        color: var(--color-text-2);
      }
      a.attr-value {
        color: rgb(var(--arcoblue-6));
      }
    }
  }
</style>
