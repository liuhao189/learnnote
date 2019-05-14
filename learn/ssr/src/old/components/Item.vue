<template>
  <div>{{item.title}}</div>
</template>
<script>
export default {
  asyncData({ route, store }) {
    return store.dispatch("fetchItem", route.params.id);
  },
  computed: {
    item() {
      return this.$store.state.items[this.$route.params.id];
    }
  }
};
</script>

