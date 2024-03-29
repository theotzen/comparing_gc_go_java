package com.javagcapp;

import java.io.IOException;
import java.lang.management.GarbageCollectorMXBean;
import java.lang.management.ManagementFactory;
import java.lang.management.MemoryMXBean;
import java.lang.management.MemoryUsage;
import java.util.List;

import io.prometheus.client.Counter;
import io.prometheus.client.Gauge;
import io.prometheus.client.exporter.HTTPServer;
import io.prometheus.client.hotspot.DefaultExports;

public class LinkedListGenerator {

    private static final Counter linkedListCounter = Counter.build()
            .name("java_linked_lists_created_total")
            .help("Total number of linked lists created.")
            .register();

    private static final Gauge heapUsed = Gauge.build()
            .name("java_heap_used_bytes")
            .help("Amount of heap memory used.")
            .register();

    private static final Gauge heapMax = Gauge.build()
            .name("java_heap_max_bytes")
            .help("Maximum amount of heap memory that can be used.")
            .register();

    private static final Gauge sys = Gauge.build()
            .name("java_sys_obtained_bytes")
            .help("Amount of heap memory committed for use.")
            .register();

    private static final Gauge heapCommitted = Gauge.build()
            .name("java_heap_committed_bytes")
            .help("Amount of heap memory committed for use.")
            .register();

    private static final Counter gcTotalTimeCounter = Counter.build()
            .name("java_gc_total_seconds")
            .help("Total seconds spent GCing")
            .register();

    private static final Counter gcCounter = Counter.build()
            .name("java_gc_count")
            .help("Number of completed GC cycles")
            .register();

    private static class Node {
        int value;
        Node next;

        Node(int value) {
            this.value = value;
            this.next = null;
        }
    }

    private static Node generateList(int size) {
        linkedListCounter.inc();
        if (size == 0) {
            return null;
        }
        Node head = new Node(0);
        Node current = head;
        for (int i = 1; i < size; i++) {
            current.next = new Node(i);
            current = current.next;
        }
        return head;
    }

    public static void main(String[] args) throws IOException {
        DefaultExports.initialize();
        HTTPServer server = new HTTPServer(2112); // Sets up Prometheus server

        new Thread(() -> {
            List<GarbageCollectorMXBean> gcBeans = ManagementFactory.getGarbageCollectorMXBeans();
            long lastGcCount = 0;
            long lastGcTime = 0;
            while (true) {
                long currentGcCount = gcBeans.stream().mapToLong(GarbageCollectorMXBean::getCollectionCount).sum();
                long gcDiff = currentGcCount - lastGcCount;
                gcCounter.inc(gcDiff);
                lastGcCount = currentGcCount;

                long currentGcTime = gcBeans.stream().mapToLong(GarbageCollectorMXBean::getCollectionTime).sum();
                long gcTimeDiff = currentGcTime - lastGcTime;
                gcTotalTimeCounter.inc(gcTimeDiff);
                lastGcTime = currentGcTime;

                MemoryMXBean memoryMXBean = ManagementFactory.getMemoryMXBean();
                MemoryUsage heapUsage = memoryMXBean.getHeapMemoryUsage();
                MemoryUsage nonHeapUsage = memoryMXBean.getNonHeapMemoryUsage();
                sys.set(heapUsage.getCommitted() + nonHeapUsage.getCommitted());
                heapUsed.set(heapUsage.getUsed());
                heapMax.set(heapUsage.getMax());
                heapCommitted.set(heapUsage.getCommitted());

                try {
                    Thread.sleep(100);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }).start();

        while (true) {
            Node list = generateList(1000000);
            try {
                Thread.sleep(100);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            list = null;
            System.gc();
        }
    }
}