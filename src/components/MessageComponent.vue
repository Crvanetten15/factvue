// src/components/MessageComponent.vue
<template>
  <div>
    <p v-if="apiMessage">API says: {{ apiMessage }}</p>
    <p v-else>Loading message...</p>
  </div>
</template>

<script>
import ApiService from '../services/ApiService';

export default {
  data() {
    return {
      apiMessage: ''
    };
  },
  created() {
    this.fetchRootMessage();
  },
  methods: {
    async fetchRootMessage() {
      try {
        const response = await ApiService.getRootMessage();
        this.apiMessage = response.data.message; // Update this line based on your API response structure
      } catch (error) {
        console.error(error);
        // Handle the error appropriately
      }
    }
  }
};
</script>
