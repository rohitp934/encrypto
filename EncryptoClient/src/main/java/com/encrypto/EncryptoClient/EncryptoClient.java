/*
 * This Java source file was generated by the Gradle 'init' task.
 */
package com.encrypto.EncryptoClient;

import static org.slf4j.LoggerFactory.getLogger;

import com.encrypto.EncryptoClient.components.LoginSignupPanel;
import com.formdev.flatlaf.themes.FlatMacDarkLaf;
import com.formdev.flatlaf.util.SystemInfo;

import net.miginfocom.swing.MigLayout;

import org.apache.log4j.BasicConfigurator;
import org.slf4j.Logger;

import java.awt.*;

import javax.swing.*;

public class EncryptoClient {
    private static final Logger logger = getLogger(EncryptoClient.class);

    public EncryptoClient() {}

    private static void render() {
        FlatMacDarkLaf.setup();
        var frame = new JFrame("Encrypto");
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.setLayout(new MigLayout());

        if (SystemInfo.isMacFullWindowContentSupported) {
            frame.getRootPane().putClientProperty("apple.awt.transparentTitleBar", true);
            frame.getRootPane().putClientProperty("apple.awt.fullWindowContent", true);
            frame.getRootPane().putClientProperty("apple.awt.transparentTitleBar", true);
            frame.getRootPane().putClientProperty("apple.awt.windowTitleVisible", false);
        }

        // Login/Signup
        var loginSignupPanel = new LoginSignupPanel();
        frame.add(loginSignupPanel, "push, grow");

        frame.pack();
        frame.setLocationRelativeTo(null);
        frame.setVisible(true);
    }

    public static void main(String[] args) {
        BasicConfigurator.configure();
        logger.info("Starting Encrypto Client");
        System.setProperty("apple.laf.useScreenMenuBar", "true");
        System.setProperty("apple.awt.application.appearance", "system");
        SwingUtilities.invokeLater(EncryptoClient::render);
    }
}
