<template>
<!-- Drawer -->
<ui-drawer viewport-height>
  <ui-drawer-header>
    <ui-drawer-title>bunker</ui-drawer-title>
  </ui-drawer-header>
  <ui-drawer-content>
    <ui-nav>
      <h1>概览</h1>
      <ui-nav-item href="/#/" active>
        <ui-icon>dashboard</ui-icon>&nbsp;仪表盘
      </ui-nav-item>
      <ui-divider></ui-divider>

      <h1>用户管理</h1>
      <ui-nav-item href="/#/users">
        <ui-icon>person</ui-icon>&nbsp;用户列表
      </ui-nav-item>
      <ui-nav-item href="/#/users2">
        <ui-icon>person</ui-icon>&nbsp;用户列表
      </ui-nav-item>
      <ui-divider></ui-divider>

      <h1>资产管理</h1>
      <ui-nav-item href="javascript:void(0)">
        <ui-icon>computer</ui-icon>&nbsp;资产列表
      </ui-nav-item>
      <ui-nav-item href="javascript:void(0)">
        <ui-icon>tag</ui-icon>&nbsp;标签管理
      </ui-nav-item>
    </ui-nav>
  </ui-drawer-content>
</ui-drawer>
</template>

<script>
export default {
  name: 'sidebar'
};
</script>

<style scoped>
a {
  text-decoration: none;
  padding-bottom: 5px;
}
.router-link-active {
  text-decoration: none;
}
</style>