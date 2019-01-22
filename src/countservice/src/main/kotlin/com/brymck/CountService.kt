package com.brymck

import gkesite.Demo.CountRequest
import gkesite.Demo.CountResponse
import gkesite.CountServiceGrpc.CountServiceImplBase
import io.grpc.Server
import io.grpc.ServerBuilder
import io.grpc.StatusRuntimeException
import io.grpc.stub.StreamObserver
import io.opencensus.contrib.grpc.metrics.RpcViews
import io.opencensus.exporter.trace.logging.LoggingTraceExporter
import org.apache.logging.log4j.Level
import org.apache.logging.log4j.LogManager

private val logger = LogManager.getLogger(CountService::class.java)

class CountService {
    companion object {
        val service = CountService()

        /** Main launches the server from the command line.  */
        @JvmStatic
        fun main(args: Array<String>) {
            // Registers all RPC views.
            RpcViews.registerAllViews()

            // Registers logging trace exporter.
            LoggingTraceExporter.register()

            // Start the RPC server. You shouldn't see any output from gRPC before this.
            logger.info("AdService starting.")
            val service = CountService.getInstance()
            service.start()
            service.blockUntilShutdown()
        }

        fun getInstance(): CountService {
            return service
        }
    }

    private var server: Server? = null

    private fun getPort(): Int {
        val port = System.getenv("PORT")
        return if (port == null) {
            9555
        } else {
            Integer.parseInt(port)
        }
    }

    fun start() {
        val port = getPort()

        server = ServerBuilder.forPort(port).addService(CountServiceImpl()).build().start()
        logger.info("Count Service started, listening on $port")
        Runtime.getRuntime().addShutdownHook(Thread {
            // Use stderr here since the logger may have been reset by its JVM shutdown hook.
            System.err.println("*** shutting down gRPC ads server since JVM is shutting down")
            server!!.shutdown()
            System.err.println("*** server shut down")
        })
    }

    internal class CountServiceImpl : CountServiceImplBase() {
        private var count = 0

        override fun getCount(request: CountRequest, responseObserver: StreamObserver<CountResponse>) {
            try {
                count++
                logger.info("Increment count to $count")
                val reply = CountResponse.newBuilder().setCount(count).build()
                responseObserver.onNext(reply)
                responseObserver.onCompleted()
            } catch (e: StatusRuntimeException) {
                logger.log(Level.WARN, "greet Failed", e.status)
            }
        }
    }

    /** Await termination on the main thread since the grpc library uses daemon threads. */
    fun blockUntilShutdown() {
        if (server != null) {
            server!!.awaitTermination()
        }
    }
}
