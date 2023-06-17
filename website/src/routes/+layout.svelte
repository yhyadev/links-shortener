<script lang="ts">
  import "../app.css";
</script>

<slot />

<style lang="postcss">
  :global(body) {
    background-color: theme(backgroundColor);
    color: theme(textColor);
  }
</style>
