<script lang="ts">
	import { Timer, Play, Pause, RotateCcw } from 'lucide-svelte';
	import { Input } from '$lib/components/ui/input';
	import { Fireworks, type FireworksOptions } from '@fireworks-js/svelte';

	// svelte-ignore non_reactive_update
	let fw: Fireworks;
	let play: boolean = $state(false);

	let options: FireworksOptions = {
		hue: { min: 0, max: 50 },
		rocketsPoint: { min: 50, max: 200 },
		brightness: { min: 0, max: 80 },
		acceleration: 1.03,
		traceSpeed: 8,
		particles: 35,
		opacity: 0.5
	};

	let timeExpired = $state(false);

	function toggleFireworks() {
		const fireworks = fw.fireworksInstance();
		if (fireworks.isRunning) {
			fireworks.waitStop();
		} else {
			fireworks.start();
		}
	}

	function padZero(num: number) {
		return num < 10 ? `0${num}` : num;
	}

	let hh = $state(padZero(new Date().getHours()));
	let mm = $state(padZero(new Date().getMinutes()));
	let ss = $state(padZero(new Date().getSeconds()));

	let timer = $state(20);
	let lastSeconds = $state(false);
	let interval: number;
	let clockInterval: number;
	let fireworkInterval: number;

	let remainingTime: number = Infinity;
	let timerStarted = $state(false);

	function startTimer(pause: boolean = false, reset: boolean = false) {
		if (reset) {
			remainingTime = Infinity;
		}

		timerStarted = true;
		timeExpired = false;

		if (interval) clearInterval(interval);
		if (clockInterval) clearInterval(clockInterval);

		let totalSeconds = timer * 60;

		if (remainingTime < totalSeconds) {
			totalSeconds = remainingTime;
		}

		if (pause) {
			return;
		}

		interval = setInterval(() => {
			if (totalSeconds <= 16) {
				lastSeconds = true;
			}
			if (totalSeconds <= 0) {
				lastSeconds = false;
				clearInterval(interval);
				timeExpired = true;
				toggleFireworks();

				return;
			}
			totalSeconds--;
			remainingTime = totalSeconds;

			const hours = Math.floor(totalSeconds / 3600);
			const minutes = Math.floor((totalSeconds % 3600) / 60);
			const seconds = totalSeconds % 60;

			hh = padZero(hours);
			mm = padZero(minutes);
			ss = padZero(seconds);
		}, 1000);

		fireworkInterval = setInterval(() => {
			if (timeExpired) {
				timeExpired = false;
				timerStarted = false;
				clearInterval(fireworkInterval);
			}
		}, 10000);
	}

	function setClockInterval() {
		return setInterval(() => {
			hh = padZero(new Date().getHours());
			mm = padZero(new Date().getMinutes());
			ss = padZero(new Date().getSeconds());
		}, 1000);
	}

	clockInterval = setClockInterval();
</script>

<div class=" h-screen text-white">
	<div class="absolute top-3 right-0 p-2 flex flex-col justify-center items-center gap-2">
		<Input
			class="pl-3 h-8 pr-0 text-xs w-10 bg-black border-0 [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
			bind:value={timer}
			type="number"
			min="1"
			max="99"
		/>
		<button
			onclick={() => startTimer(false, true)}
			class="shadow-lg p-2 rounded-full bg-black text-white hover:bg-slate-800"
		>
			{#if timerStarted}
				<RotateCcw class="h-5 w-5" />
			{:else}
				<Timer class="h-5 w-5" />
			{/if}
		</button>
	</div>
	<div class="p-4 flex items-center w-full h-full">
		<div class="flex gap-5">
			<div>
				<span class="text-7xl font-bold">{hh}</span>
			</div>
			<div>
				<span class="text-7xl font-bold">:</span>
			</div>
			<div>
				<span class="text-7xl font-bold">{mm}</span>
			</div>
			<div class="flex flex-col justify-between items-center">
				<button
					onclick={() => {
						play = !play;
						startTimer(play);
					}}
					class="p-2 rounded-full bg-black text-white hover:bg-slate-800 {timerStarted
						? ''
						: 'invisible'}"
				>
					{#if play}
						<Play class="h-5 w-5" />
					{:else}
						<Pause class="h-5 w-5" />
					{/if}
				</button>

				<!-- <Pause class="h-8 w-8" /> -->
				<span class="text-2xl font-bold {lastSeconds ? 'text-red-400' : ''}">{ss}</span>
			</div>
		</div>
	</div>
	{#if timeExpired}
		<Fireworks bind:this={fw} {options} class="absolute top-0 h-full z-10" />
	{/if}
</div>
