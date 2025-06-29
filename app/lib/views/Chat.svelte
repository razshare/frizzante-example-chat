<style>
    .chat {
        text-align: start;
    }
    .message {
        width: 100%;
    }
</style>

<script lang="ts">
    import Layout from "$lib/components/Layout.svelte"
    import { source } from "$frizzante/scripts/source.ts"
    import { action } from "$frizzante/scripts/action.ts"

    type Props = { messages: string[] }
    let { messages }: Props = $props()
    const connection = source("/chat/messages/stream")
    const channel = connection.select()
    channel.subscribe(message => messages.push(message))
</script>

<Layout title="Chat Room">
    <h1>Chat Room</h1>
    <div class="chat">
        {#each messages as message, id (id)}
            <div>{message}</div>
        {/each}
    </div>
    <br />
    <form method="POST" {...action("/chat/messages/add")}>
        <input class="message" type="text" name="message" />
        <br />
        <br />
        <button class="link">Submit Message</button>
    </form>
</Layout>
