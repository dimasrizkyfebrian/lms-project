<template>
  <div class="flex h-screen bg-gray-100 dark:bg-gray-900">
    <AppSidebar />

    <div class="flex-1 flex flex-col overflow-hidden">
      <header class="bg-white dark:bg-gray-800 shadow-sm p-4 flex justify-end">
        <UButton color="red" variant="soft" @click="handleLogout">
          Log Out
        </UButton>
      </header>

      <main
        class="flex-1 overflow-x-hidden overflow-y-auto bg-gray-100 dark:bg-gray-900 p-8"
      >
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup>
const token = useCookie("token");
const user = useUser();

const handleLogout = async () => {
  token.value = null;
  user.value = null;
  await navigateTo("/login");
};
</script>
