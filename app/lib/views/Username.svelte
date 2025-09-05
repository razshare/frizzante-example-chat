<script lang="ts">
    import Layout from "$lib/components/Layout.svelte"
    import {action} from "$lib/scripts/core/action.ts";
</script>

<Layout title="Pick Room">
    <h1>Pick a valid username.</h1>
    <form method="POST" {...action("/username")}>
        <input type="text" name="username" />
        <button class="link">Submit Username</button>
    </form>
</Layout>
