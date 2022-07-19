<template>
<ui-drawer viewport-height>
  <ui-drawer-header>
    <ui-drawer-title>bunker</ui-drawer-title>
  </ui-drawer-header>
  <ui-drawer-content>
    <ui-list-group>
      <div v-for="(group,i) in groups">
        <ui-list-group-subheader><h3>{{group.name}}</h3></ui-list-group-subheader>
        <router-link
          v-for="page in group.pages"
          v-slot="{href,isActive}"
          :to="page.path"
          style="text-decoration:none;">
          <ui-nav-item :href="href" :active="isActive">
            <ui-item-first-content>
              <ui-icon>{{page.icon}}</ui-icon>
            </ui-item-first-content>
            <ui-item-text-content>{{page.name}}</ui-item-text-content>
          </ui-nav-item>
          </router-link>
        <ui-list-divider v-if="i!==groups.length-1"></ui-list-divider>
      </div>
    </ui-list-group>
  </ui-drawer-content>
</ui-drawer>
</template>

<script>
import { groups } from '@/router';

export default {
  name: 'sidebar',
  setup() {
    return {
      groups: groups
    };
  }
};
</script>

<style scoped>
</style>