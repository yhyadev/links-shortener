<script lang="ts">
	import axios from 'axios';
	import { onMount } from 'svelte';
	import { PUBLIC_API_URL } from '$env/static/public';

	export let inputurl = '';
	export let finalurl = '';

	function shortenURL(url: string) {
		let response = axios.post(PUBLIC_API_URL + '/shorten', {
			url
		});

		return response;
	}

	onMount(() => {
		let form = document.getElementById('shorten-form')!;

		form.onsubmit = (e) => {
			e.preventDefault();

			shortenURL(inputurl)
				.then((res) => {
					finalurl = PUBLIC_API_URL + '/' + res.data.slug;

					document.getElementById('error')!.innerText = '';
					document.getElementById('url')!.classList.remove('input-error');
				})
				.catch(() => {
					document.getElementById('error')!.innerText = 'Invalid URL';
					document.getElementById('url')!.classList.add('input-error');
				});
		};
	});
</script>

<svelte:head>
	<title>Links Shortener</title>
</svelte:head>

<div class="hero lg:hero-xl min-h-screen">
	<form class="form-control" id="shorten-form">
		<input
			type="text"
			id="url"
			placeholder="URL"
			class="input input-bordered input-primary w-full"
			bind:value={inputurl}
			required
		/>
		<p id="error" class="label label-text-alt" style="color: red;" />
		<input type="submit" value="Shorten" class="btn btn-primary btn-wide" />
	</form>
	{#if finalurl}
		<div class="card" style="background-color: #404040;">
			<div class="card-body">
				<h1 class="card-title">Shortened</h1>
				<a href={finalurl} class="text-primary underline">{finalurl}</a>
				<div class="card-actions">
					<input
						type="button"
						value="Copy Link!"
						class="btn btn-primary"
						on:click={(e) => {
							navigator.clipboard.writeText(finalurl);
							e.currentTarget.value = 'Copied!';
						}}
					/>
				</div>
			</div>
		</div>
	{/if}
</div>
