<script lang="ts">
    import Layout from "$lib/components/Layout.svelte"
    import { action } from "$frizzante/core/scripts/action.ts"
</script>

<Layout title="Pick Room">
    <h1>Pick a valid username.</h1>
    <form method="POST" {...action("/chat/username/set")}>
        <input type="text" name="username" />
        <button class="link">Submit Username</button>
    </form>
</Layout>
