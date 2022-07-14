<template>
<ui-drawer viewport-height>
  <ui-drawer-header>
    <ui-drawer-title>bunker</ui-drawer-title>
  </ui-drawer-header>
  <ui-drawer-content>
    <ui-list-group>
      <ui-list-group-subheader>概览</ui-list-group-subheader>
      <ui-nav>
        <router-link
          v-slot="{ href, isActive }"
          to="/"
          style="text-decoration:none;">
          <ui-nav-item :href="href" :active="isActive">
            <ui-item-first-content>
              <ui-icon>dashboard</ui-icon>
            </ui-item-first-content>
            <ui-item-text-content>仪表盘</ui-item-text-content>
          </ui-nav-item>
        </router-link>
      </ui-nav>

      <ui-list-divider></ui-list-divider>

      <ui-list-group-subheader>用户管理</ui-list-group-subheader>
      <ui-nav>
        <router-link
          v-slot="{ href, isActive }"
          to="/users"
          style="text-decoration:none;">
          <ui-nav-item :href="href" :active="isActive">
            <ui-item-first-content>
              <ui-icon>person</ui-icon>
            </ui-item-first-content>
            <ui-item-text-content>用户列表</ui-item-text-content>
          </ui-nav-item>
        </router-link>
        <router-link
          v-slot="{ href, isActive }"
          to="/users2"
          style="text-decoration:none;">
          <ui-nav-item :href="href" :active="isActive">
            <ui-item-first-content>
              <ui-icon>person</ui-icon>
            </ui-item-first-content>
            <ui-item-text-content>用户列表</ui-item-text-content>
          </ui-nav-item>
        </router-link>
      </ui-nav>
      <ui-list-divider></ui-list-divider>

      <ui-list-group-subheader>资产管理</ui-list-group-subheader>
      <ui-nav>
        <ui-nav-item href="javascript:void(0)">
          <ui-icon>computer</ui-icon>&nbsp;资产列表
        </ui-nav-item>
        <ui-nav-item href="javascript:void(0)">
          <ui-icon>tag</ui-icon>&nbsp;标签管理
        </ui-nav-item>
      </ui-nav>
    </ui-list-group>
  </ui-drawer-content>
</ui-drawer>
</template>

<script>
export default {
  name: 'sidebar'
};
</script>

<style scoped>
</style>