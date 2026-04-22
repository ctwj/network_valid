<template>
  <router-view v-slot="{ Component, route }">
    <transition name="fade" mode="out-in">
      <keep-alive v-if="route.meta?.keepAlive">
        <component :is="Component" :key="route.fullPath" />
      </keep-alive>
      <component :is="Component" :key="route.fullPath" v-else />
    </transition>
  </router-view>
</template>

<script lang="ts">
  export default {
    name: 'MainView',
  };
</script>